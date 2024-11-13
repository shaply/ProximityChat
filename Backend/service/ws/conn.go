package conn

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/shaply/ProximityChat/Backend/service/auth"
	"github.com/shaply/ProximityChat/Backend/types"
)

type Handler struct {
	store *types.UserStore
}

type Client struct {
	Conn  *websocket.Conn
	Email string
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// Security flaw: Doesn't check the origin of the request
	CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make(map[*Client]bool)

func NewHandler(store *types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/ws/{JWTToken}", auth.WithJWTAuth(serveWS, *h.store)).Methods("GET")
}

// serveWS upgrades the connection to a websocket connection
func serveWS(w http.ResponseWriter, r *http.Request) {
	// Upgrade the connection to a websocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "could not upgrade connection", http.StatusBadRequest)
		return
	}

	client := &Client{
		Conn:  conn,
		Email: r.Context().Value(auth.EmailKey).(string),
	}

	// Register the client
	clients[client] = true

	fmt.Println("Client connected:", client.Email)

	// Listen for messages
	// TODO
}
