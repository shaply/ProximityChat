package quadtree

type Circle struct {
	Center Point
	Radius int
}

func NewCircle(center Point, radius int) Circle {
	return Circle{Center: center, Radius: radius}
}

func (c *Circle) Contains(point Point) bool {
	return (point.X-c.Center.X)*(point.X-c.Center.X)+(point.Y-c.Center.Y)*(point.Y-c.Center.Y) <= c.Radius*c.Radius
}

// This checks if the point would be in the circle at the origin with the radius of the current circle.
func (c *Circle) ContainsFromOriginWithRadius(point Point) bool {
	return point.X*point.X+point.Y*point.Y <= c.Radius*c.Radius
}
