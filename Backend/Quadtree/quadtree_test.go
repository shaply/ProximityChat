package quadtree

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
	q.Insert(Point{20, 20})
	q.Insert(Point{30, 30})
	q.Insert(Point{40, 40})
	q.Insert(Point{50, 50})
	q.Insert(Point{60, 60})
	q.Insert(Point{70, 70})
	q.Insert(Point{80, 80})
	q.Insert(Point{90, 90})
	q.Insert(Point{100, 100})

	q.RemovePoint(Point{20, 20})
	q.RemovePoint(Point{70, 70})
	q.RemovePoint(Point{30, 31})

	fmt.Println(q.String())
}
