package quadtree

/*
TODO for quadtree: Test QueryNearby
TODO for bounds: Test
*/

import (
	"context"
	"fmt"
	"testing"

	"github.com/shaply/ProximityChat/Backend/types"
)

type test struct {
	name string
}

var rTest = &test{"rand"}

func TestInsertAndPrint(t *testing.T) {
	q := NewQuadtree(NewBounds(NewPoint(0, 0), NewPoint(100, 100)), nil)
	q.Insert(NewPointWithClient(10, 10, rTest))
	q.Insert(NewPointWithClient(20, 20, rTest))
	q.Insert(NewPointWithClient(30, 30, rTest))
	q.Insert(NewPointWithClient(40, 40, rTest))
	q.Insert(NewPointWithClient(50, 50, rTest))
	q.Insert(NewPointWithClient(60, 60, rTest))
	q.Insert(NewPointWithClient(70, 70, rTest))
	q.Insert(NewPointWithClient(80, 80, rTest))
	q.Insert(NewPointWithClient(90, 90, rTest))
	q.Insert(NewPointWithClient(100, 100, rTest))

	fmt.Println(q.String())

}

func TestRemove(t *testing.T) {
	q := NewQuadtree(NewBounds(NewPoint(0, 0), NewPoint(100, 100)), nil)
	q.Insert(NewPointWithClient(10, 10, rTest))
	q.Insert(NewPointWithClient(20, 20, rTest))
	q.Insert(NewPointWithClient(30, 30, rTest))
	q.Insert(NewPointWithClient(40, 40, rTest))
	q.Insert(NewPointWithClient(50, 50, rTest))
	q.Insert(NewPointWithClient(60, 60, rTest))
	q.Insert(NewPointWithClient(70, 70, rTest))
	q.Insert(NewPointWithClient(80, 80, rTest))
	q.Insert(NewPointWithClient(90, 90, rTest))
	q.Insert(NewPointWithClient(100, 100, rTest))
	q.Insert(NewPointWithClient(0, 0, rTest))

	q.RemovePoint(NewPointWithClient(20, 20, rTest))
	q.RemovePoint(NewPointWithClient(70, 70, rTest))
	q.RemovePoint(NewPointWithClient(30, 31, rTest))
	q.RemovePoint(NewPointWithClient(0, 0, rTest))
	q.RemovePoint(NewPointWithClient(50, 50, rTest))

	fmt.Println(q.String())
}

func TestMovePoint(t *testing.T) {
	q := NewQuadtree(NewBounds(NewPoint(0, 0), NewPoint(100, 100)), nil)
	q.Insert(NewPointWithClient(10, 10, rTest))
	q.Insert(NewPointWithClient(20, 20, rTest))
	q.Insert(NewPointWithClient(30, 30, rTest))
	q.Insert(NewPointWithClient(40, 40, rTest))
	q.Insert(NewPointWithClient(50, 50, rTest))
	q.Insert(NewPointWithClient(60, 60, rTest))
	q.Insert(NewPointWithClient(70, 70, rTest))
	q.Insert(NewPointWithClient(80, 80, rTest))
	q.Insert(NewPointWithClient(90, 90, rTest))
	q.Insert(NewPointWithClient(100, 100, rTest))
	q.Insert(NewPointWithClient(100, 100, rTest))

	q.MovePoint(NewPointWithClient(0, 0, rTest), NewPointWithClient(25, 75, rTest))
	q.MovePoint(NewPointWithClient(10, 10, rTest), NewPointWithClient(25, 25, rTest))
	q.MovePoint(NewPointWithClient(50, 50, rTest), NewPointWithClient(75, 25, rTest))
	q.MovePoint(NewPointWithClient(91, 90, rTest), NewPointWithClient(25, 35, rTest))

	fmt.Println(q)
}

func TestCircleIntersectBoundary(t *testing.T) {
	q := NewQuadtree(NewBounds(NewPoint(0, 0), NewPoint(100, 100)), nil)
	b := NewBounds(NewPoint(10, 35), NewPoint(30, 75))
	c := NewCircle(*NewPoint(90, 50), 25)

	bb := q.Bounds.WrapIntersectionBoundCircle(b, c)
	fmt.Println(bb)
}

func TestQueryNearby(t *testing.T) {
	q := NewQuadtree(NewBounds(NewPoint(0, 0), NewPoint(100, 100)), nil)
	q.Insert(NewPointWithClient(10, 10, rTest))
	q.Insert(NewPointWithClient(20, 20, rTest))
	q.Insert(NewPointWithClient(30, 30, rTest))
	q.Insert(NewPointWithClient(40, 40, rTest))
	q.Insert(NewPointWithClient(50, 50, rTest))
	q.Insert(NewPointWithClient(60, 60, rTest))
	q.Insert(NewPointWithClient(70, 70, rTest))
	q.Insert(NewPointWithClient(80, 80, rTest))
	q.Insert(NewPointWithClient(90, 90, rTest))
	q.Insert(NewPointWithClient(100, 100, rTest))
	q.Insert(NewPointWithClient(65, 65, rTest))
	q.Insert(NewPointWithClient(50, 50, rTest))

	query := q.QueryNearby(NewPoint(50, 50), 25)
	fmt.Println("Query: " + query.String())

	fmt.Println("test 2")

	q.MovePoint(NewPointWithClient(50, 50, rTest), NewPointWithClient(25, 25, rTest))
	query = q.QueryNearby(NewPoint(90, 90), 25)
	fmt.Println("Query: " + query.String())
}

func TestQuadHandler(t *testing.T) {
	client1 := types.Client{Email: "test1", Location: types.Location{Lat: 0, Lon: 0}}
	client2 := types.Client{Email: "test2", Location: types.Location{Lat: 0, Lon: 0}}

	var h = QuadHandler
	h.Insert(&client1)
	h.Insert(&client2)

	h.PrintTree()

	oldLoc := client1.Location
	client1.Location = types.Location{Lat: 10, Lon: 10}
	h.Move(&client1, &oldLoc)

	h.PrintTree()

	fmt.Println("Nearby clients:")
	ch := h.GetNearby(context.Background(), &client1, 25)
	for c := range ch {
		fmt.Println(c.Email)
	}
}

func TestRandom(t *testing.T) {
	type test struct {
		name string
	}
	f := &test{"rand"}
	g := &test{"rand"}
	fmt.Println(f == g)
}
