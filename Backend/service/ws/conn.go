package ws

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/shaply/ProximityChat/Backend/service/auth"
	quadtree "github.com/shaply/ProximityChat/Backend/service/ws/quadtree"
	"github.com/shaply/ProximityChat/Backend/types"
)

type Handler struct {
	store types.UserStore
	qTree *quadtree.Handler // This will be used to manage all the connections
}

type EnhancedMessage struct {
	Type      string    `json:"type"`
	Text      string    `json:"text,omitempty"`
	Email     string    `json:"email"`
	Timestamp time.Time `json:"timestamp"`
}

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,

		// Security flaw: Doesn't check the origin of the request
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	dist            = 25.0 // Distance in miles
	ClientList      = make(map[*types.Client]chan types.Message)
	ClientListMutex sync.Mutex
	ConnHandler     = NewHandler(nil)
)

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store, qTree: quadtree.QuadHandler}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/ws/{JWTToken}", auth.WithJWTAuth(serveWS, h.store)).Methods("GET")
}

// serveWS upgrades the connection to a websocket connection
func serveWS(w http.ResponseWriter, r *http.Request) {
	// Upgrade the connection to a websocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		http.Error(w, "could not upgrade connection", http.StatusBadRequest)
		return
	}
	defer conn.Close()

	client := NewClient(conn, r.Context().Value(auth.EmailKey).(string), []float64{0, 0})

	// Register the client
	ClientList[client] = make(chan types.Message, 100) // Only 100 messages allowed at a time
	ConnHandler.qTree.Insert(client)

	fmt.Println("Client connected:", client.Email)
	ConnHandler.qTree.PrintTree()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Listen for messages
	go ConnHandler.HandleMessages(ctx, client)
	readMessages(ctx, client) // Don't want to make this a go routine because if this function returns, the connection is closed
}

func (h *Handler) HandleMessages(ctx context.Context, client *types.Client) {
	defer func() {
		log.Printf("Closing message handler for client: %s\n", client.Email)
	}()
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-ClientList[client]:
			// Figure out whether message is a location or text
			if msg.Type == "location" {
				// Update the client's location
				oldLoc := client.Location
				UpdateLocation(client, msg.Location)
				ConnHandler.qTree.Move(client, &oldLoc)
				sucess := types.Message{
					Type:    "status",
					Message: "Location updated",
				}
				client.Conn.WriteJSON(sucess)
			} else if msg.Type == "text" {
				// Broadcast the message to all clients
				enhancedMsg := EnhancedMessage{
					Type:      "text",
					Text:      msg.Message,
					Email:     client.Email,
					Timestamp: time.Now(),
				}
				ch := ConnHandler.qTree.GetNearby(ctx, client, dist)
			broadcastLoop:
				for {
					select {
					case <-ctx.Done():
						return
					case c, ok := <-ch:
						if !ok {
							break broadcastLoop
						}
						c.Conn.WriteJSON(enhancedMsg)
					}
				}
			}

			ConnHandler.qTree.PrintTree()
		}
	}
}

// readMessages reads messages from the client and broadcasts them
func readMessages(ctx context.Context, client *types.Client) {
	log.Printf("Reading messages from client: %s\n", client.Email)
	defer func() {
		client.Conn.Close()
		ClientListMutex.Lock()
		delete(ClientList, client)
		ConnHandler.qTree.Remove(client)
		ClientListMutex.Unlock()
		log.Printf("Client disconnected: %s\n", client.Email)
	}()

	for {
		var msg types.Message
		err := client.Conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}
		select {
		case <-ctx.Done():
			return
		case ClientList[client] <- msg:
		}
	}
}
