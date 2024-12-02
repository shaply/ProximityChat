package quadtree

type Bounds struct {
	TopLeft     Point
	BottomRight Point
}

func NewBounds(topLeft Point, bottomRight Point) Bounds {
	return Bounds{
		TopLeft:     topLeft,
		BottomRight: bottomRight,
	}
}
