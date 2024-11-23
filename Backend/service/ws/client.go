package ws

import (
	"github.com/gorilla/websocket"
	"github.com/shaply/ProximityChat/Backend/types"
)

func NewClient(conn *websocket.Conn, email string) *types.Client {
	return &types.Client{
		Conn:  conn,
		Email: email,
	}
}
