package quadtree

import "math"

type Bounds struct {
	BottomLeft Point
	TopRight   Point
}

func NewBounds(point1 Point, point2 Point) Bounds {
	return Bounds{
		BottomLeft: Point{X: int(math.Min(float64(point1.X), float64(point2.X))), Y: int(math.Min(float64(point1.Y), float64(point2.Y)))},
		TopRight:   Point{X: int(math.Max(float64(point1.X), float64(point2.X))), Y: int(math.Max(float64(point1.Y), float64(point2.Y)))},
	}
}

func NewBoundsWithPointArray(points []Point) *Bounds {
	if len(points) == 0 {
		return nil
	}
	bounds := NewBounds(points[0], points[0])
	for _, p := range points {
		bounds.Extend(p)
	}
	return &bounds
}

func (b *Bounds) Width() int {
	return b.TopRight.X - b.BottomLeft.X
}

func (b *Bounds) Height() int {
	return b.TopRight.Y - b.BottomLeft.Y
}

func (b *Bounds) Contains(point Point) bool {
	return point.X >= b.BottomLeft.X && point.X <= b.TopRight.X && point.Y >= b.BottomLeft.Y && point.Y <= b.TopRight.Y
}

func (b *Bounds) Area() int {
	return b.Width() * b.Height()
}

func (b *Bounds) SplitInto4() *[4]Bounds {
	return &[4]Bounds{
		NewBounds(Point{b.BottomLeft.X + b.Width()/2, b.BottomLeft.Y + b.Height()/2}, b.TopRight),
		NewBounds(Point{b.BottomLeft.X, b.BottomLeft.Y + b.Height()/2}, Point{b.BottomLeft.X + b.Width()/2, b.TopRight.Y}),
		NewBounds(b.BottomLeft, Point{b.BottomLeft.X + b.Width()/2, b.BottomLeft.Y + b.Height()/2}),
		NewBounds(Point{b.BottomLeft.X + b.Width()/2, b.BottomLeft.Y}, Point{b.TopRight.X, b.BottomLeft.Y + b.Height()/2}),
	}
}

// Creates a new boundary to accomodate the point
func (b *Bounds) Extend(point Point) {
	if point.X < b.BottomLeft.X {
		b.BottomLeft.X = point.X
	}
	if point.Y < b.BottomLeft.Y {
		b.BottomLeft.Y = point.Y
	}
	if point.X > b.TopRight.X {
		b.TopRight.X = point.X
	}
	if point.Y > b.TopRight.Y {
		b.TopRight.Y = point.Y
	}
}

/**
 * Returns the quadrant in which the point is located
 * @Returns -1 if the point is not in the bounds, otherwise, 0, 1, 2, or 3
 */
func (b *Bounds) WhichQuadrant(point Point) int8 {
	if !b.Contains(point) {
		return -1
	}
	if point.X >= b.BottomLeft.X+b.Width()/2 {
		if point.Y >= b.BottomLeft.Y+b.Height()/2 {
			return 0
		}
		return 3
	} else {
		if point.Y >= b.BottomLeft.Y+b.Height()/2 {
			return 1
		}
		return 2
	}
}
