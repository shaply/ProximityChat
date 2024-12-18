package quadtree

/*
TODO for quadtree: Test QueryNearby
TODO for bounds: Test
*/

import (
	"fmt"
	"testing"
)

func TestInsertAndPrint(t *testing.T) {
	q := NewQuadtree(NewBounds(Point{0, 0}, Point{100, 100}), nil)
	q.Insert(Point{10, 10})
	q.Insert(Point{20, 20})
	q.Insert(Point{30, 30})
	q.Insert(Point{40, 40})
	q.Insert(Point{50, 50})
	q.Insert(Point{60, 60})
	q.Insert(Point{70, 70})
	q.Insert(Point{80, 80})
	q.Insert(Point{90, 90})
	q.Insert(Point{100, 100})

	fmt.Println(q.String())

}

func TestRemove(t *testing.T) {
	q := NewQuadtree(NewBounds(Point{0, 0}, Point{100, 100}), nil)
	q.Insert(Point{10, 10})
	q.Insert(Point{20, 20})
	q.Insert(Point{30, 30})
	q.Insert(Point{40, 40})
	q.Insert(Point{50, 50})
	q.Insert(Point{60, 60})
	q.Insert(Point{70, 70})
	q.Insert(Point{80, 80})
	q.Insert(Point{90, 90})
	q.Insert(Point{100, 100})
	q.Insert(Point{0, 0})

	q.RemovePoint(Point{20, 20})
	q.RemovePoint(Point{70, 70})
	q.RemovePoint(Point{30, 31})
	q.RemovePoint(Point{100, 100})
	q.RemovePoint(Point{50, 50})

	fmt.Println(q.String())
}

func TestMovePoint(t *testing.T) {
	q := NewQuadtree(NewBounds(Point{0, 0}, Point{100, 100}), nil)
	q.Insert(Point{10, 10})
	q.Insert(Point{20, 20})
	q.Insert(Point{30, 30})
	q.Insert(Point{40, 40})
	q.Insert(Point{50, 50})
	q.Insert(Point{60, 60})
	q.Insert(Point{70, 70})
	q.Insert(Point{80, 80})
	q.Insert(Point{90, 90})
	q.Insert(Point{100, 100})
	q.Insert(Point{100, 100})

	q.MovePoint(Point{100, 100}, Point{25, 75})
	q.MovePoint(Point{10, 10}, Point{25, 25})
	q.MovePoint(Point{50, 50}, Point{75, 25})
	q.MovePoint(Point{91, 90}, Point{25, 35})

	fmt.Println(q)
}

func TestCircleIntersectBoundary(t *testing.T) {
	b := NewBounds(Point{65, 65}, Point{75, 75})
	c := NewCircle(Point{50, 50}, 25)

	bb := b.IntersectsCircle(c)
	fmt.Println(bb)
}

func TestQueryNearby(t *testing.T) {
	q := NewQuadtree(NewBounds(Point{0, 0}, Point{100, 100}), nil)
	q.Insert(Point{10, 10})
	q.Insert(Point{20, 20})
	q.Insert(Point{30, 30})
	q.Insert(Point{40, 40})
	q.Insert(Point{50, 50})
	q.Insert(Point{60, 60})
	q.Insert(Point{70, 70})
	q.Insert(Point{80, 80})
	q.Insert(Point{90, 90})
	q.Insert(Point{100, 100})
	q.Insert(Point{65, 65})
	q.Insert(Point{50, 50})

	query := q.QueryNearby(Point{50, 50}, 25)
	for point := range query.Range() {
		fmt.Println(point)
	}

	fmt.Println("test 2")

	q.MovePoint(Point{50, 50}, Point{25, 25})
	query = q.QueryNearby(Point{50, 50}, 25)
	for point := range query.Range() {
		fmt.Println(point)
	}

	fmt.Println(-5 % 4)
}
