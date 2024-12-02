package quadtree

type Quadtree struct {
	Bounds Bounds

	Parent   *Quadtree
	Children [4]*Quadtree
}

func NewQuadtree(bounds Bounds) *Quadtree {
	return &Quadtree{
		Bounds: bounds,
	}
}

func (q *Quadtree) Insert(point Point) {
	// implement
}

func (q *Quadtree) Remove(point Point) {
	// implement
}

func (q *Quadtree) QueryRange(rangeBounds Bounds) []Point {
	// implement
	return nil
}

func (q *Quadtree) QueryPoint(point Point) *Point {
	// implement
	return nil
}

func (q *Quadtree) QueryNearby(point Point, dist int) []Point {
	// implement
	return nil
}
