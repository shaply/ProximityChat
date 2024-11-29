package ws

import (
	"github.com/gorilla/websocket"
	"github.com/shaply/ProximityChat/Backend/types"
)

func NewClient(conn *websocket.Conn, email string, location []float64) *types.Client {
	return &types.Client{
		Conn:  conn,
		Email: email,
		Location: types.Location{
			Lat: location[0],
			Lon: location[1],
		},
	}
}

func UpdateLocation(c *types.Client, loc []float64) {
	c.Location.Lat = loc[0]
	c.Location.Lon = loc[1]
}

func CheckDistance(c1, c2 *types.Client, distance float64) bool { // NEED TO FIX
	return c1.Location.Lat-c2.Location.Lat <= distance && c1.Location.Lon-c2.Location.Lon <= distance
}
