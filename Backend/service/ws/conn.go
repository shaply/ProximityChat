package ws

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/shaply/ProximityChat/Backend/service/auth"
	"github.com/shaply/ProximityChat/Backend/types"
)

type Handler struct {
	store types.UserStore
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// Security flaw: Doesn't check the origin of the request
	CheckOrigin: func(r *http.Request) bool { return true },
}

var ClientList = make(map[*types.Client]bool)
var ClientListMutex sync.Mutex
var Broadcast = make(chan types.Message)

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
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

	client := NewClient(conn, r.Context().Value(auth.EmailKey).(string))

	// Register the client
	ClientList[client] = true

	fmt.Println("Client connected:", client.Email)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Listen for messages
	go readMessages(ctx, client)
}

func (h *Handler) HandleMessages() {
	for {
		msg := <-Broadcast
		for client := range ClientList {
			err := client.Conn.WriteJSON(msg)
			if err != nil {
				fmt.Println("Error sending message:", err)
				client.Conn.Close()
				delete(ClientList, client)
			}
		}
	}
}

// readMessages reads messages from the client and broadcasts them
func readMessages(ctx context.Context, client *types.Client) {
	defer func() {
		client.Conn.Close()
		ClientListMutex.Lock()
		delete(ClientList, client)
		ClientListMutex.Unlock()
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
		case Broadcast <- msg:
		}
	}
}
