package quadtree

type Point struct {
	X int
	Y int
}

func Translate(p Point, x, y int) Point {
	return Point{X: p.X + x, Y: p.Y + y}
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
