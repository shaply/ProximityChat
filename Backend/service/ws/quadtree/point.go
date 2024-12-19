package quadtree

import (
	"fmt"
)

type Point struct {
	X      int
	Y      int
	Client interface{}
}

func NewPoint(x, y int) *Point {
	return &Point{X: x, Y: y}
}

func NewPointWithClient(x, y int, client interface{}) *Point {
	return &Point{X: x, Y: y, Client: client}
}

func Translate(p *Point, x, y int) *Point {
	return &Point{X: p.X + x, Y: p.Y + y}
}

func (p *Point) Translate(x, y int) {
	p.X += x
	p.Y += y
}

func (p *Point) Teleport(x, y int) {
	p.X = x
	p.Y = y
}

func CopyPoint(p *Point) Point {
	return Point{X: p.X, Y: p.Y}
}

func (p *Point) Equals(other *Point) bool {
	return p.X == other.X && p.Y == other.Y && p.Client == other.Client
}

func (p *Point) String() string {
	return fmt.Sprintf("(%v, %v, %v)", p.X, p.Y, p.Client)
}
