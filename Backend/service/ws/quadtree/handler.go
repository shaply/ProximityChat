// Connects the quadtree with the rest of the codebase.
package quadtree

import (
	"context"
	"fmt"

	"github.com/shaply/ProximityChat/Backend/types"
)

type Handler struct {
	quadtree *Quadtree
}

var QuadHandler = newHandler(0, 0, 1000, 1000)

func newHandler(x1, y1, x2, y2 int) *Handler {
	return &Handler{quadtree: NewQuadtreeWithNumbers(x1, y1, x2, y2)}
}

// New client connection
func (h *Handler) Insert(client *types.Client) {
	p := h.clientToPoint(client)
	h.quadtree.Insert(p)
}

// Client connection is gone
func (h *Handler) Remove(client *types.Client) {
	p := h.clientToPoint(client)
	h.quadtree.RemovePoint(p)
}

// Client has moved
func (h *Handler) Move(client *types.Client, oldLocation *types.Location) {
	x, y := h.locToCoord(oldLocation)
	p := NewPointWithClient(x, y, client)
	h.quadtree.TeleportPoint(p, h.clientToPoint(client))
}

// Distance is in miles
func (h *Handler) GetNearby(ctx context.Context, client *types.Client, dist float64) <-chan *types.Client {
	p := h.clientToPoint(client)
	convDist := h.coorToTree(dist)
	points := h.quadtree.QueryNearby(p, convDist)
	ch := make(chan *types.Client)
	go func() {
		defer func() {
			close(ch)
			fmt.Println("Closed channel")
		}()

		for it := points.Iterator(); it.HasNext(); {
			select {
			case <-ctx.Done():
				return
			case ch <- it.Next().Client.(*types.Client):
			}
		}
	}()
	return ch
}

func (h *Handler) PrintTree() {
	fmt.Println(h.quadtree.String())
}

func (h *Handler) locToCoord(Location *types.Location) (int, int) { // TO FIX
	p := NewPoint(int(Location.Lat), int(Location.Lon))
	h.quadtree.Bounds.WrapMovePoint(p, 0, 0, true)
	return p.X, p.Y
}

func (h *Handler) clientToPoint(client *types.Client) *Point {
	x, y := h.locToCoord(&client.Location)
	return NewPointWithClient(x, y, client)
}

func (h *Handler) coorToTree(coor float64) int { // TO FIX
	return int(coor)
}
