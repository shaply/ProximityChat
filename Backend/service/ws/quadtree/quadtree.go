// How the quadtree will work:
// 1. For the query, the quadtree will check all the quadrants that intersect with the circle
// 2. Each node in the quadtree will have a certain capacity for how many points it can store
// 3. If the capacity is reached, the node will split into 4 children
// 4. There will be a minimum size for how much the node can shrink
// 5. Removals that cause the node to go below the minimum size will cause the node to merge with its parent
// 6. Because of how bounds is working, the boundary of the big box will probably have to be a power of 2

// TODO: queryNearby implementation change to calculate with wrapping

package quadtree

import (
	"fmt"
	"strings"
)

var (
	// The maximum number of points a node can store before splitting
	MaxPoints = 4
	// The minimum area of a node before it can't split anymore
	MinArea = 5
)

// The quadrants that are split up into will go Q1 Q2 Q3 Q4
type Quadtree struct {
	Root   *Quadtree // The top parent of the entire tree
	Bounds Bounds

	IsLeaf      bool
	Parent      *Quadtree
	Points      []Point
	TotalPoints int
	Children    [4]*Quadtree
}

// For creating a completely new quadtree
func NewQuadtree(bounds Bounds, parent *Quadtree) *Quadtree {
	q := &Quadtree{
		Root:        nil,
		Bounds:      bounds,
		TotalPoints: 0,
		Points:      make([]Point, 0),
		Children:    [4]*Quadtree{nil, nil, nil, nil},
		Parent:      parent,
		IsLeaf:      true,
	}
	q.Root = q
	return q
}

// For creating a subtree in an existing tree
func NewQuadtreeWithRoot(bounds Bounds, parent *Quadtree, root *Quadtree) *Quadtree {
	q := &Quadtree{
		Root:        root,
		Bounds:      bounds,
		TotalPoints: 0,
		Points:      make([]Point, 0),
		Children:    [4]*Quadtree{nil, nil, nil, nil},
		Parent:      parent,
		IsLeaf:      true,
	}
	return q
}

// Inserts a point into the quadtree, recursively.
// Will update the trees to the quadtree that initially called the function.
func (q *Quadtree) Insert(point Point) {
	if !q.Bounds.Contains(point) {
		panic("Point is out of bounds")
	}

	if (q.IsLeaf && q.TotalPoints < MaxPoints) || q.Bounds.Width() <= MinArea {
		q.Points = append(q.Points, point)
		q.TotalPoints++
		return
	} else if q.IsLeaf {
		q.IsLeaf = false
		// Split the node into 4 children
		for i, bounds := range q.Bounds.SplitInto4() {
			q.Children[i] = NewQuadtreeWithRoot(bounds, q, q.Root)
		}
		// Move all the points into the children
		for _, p := range q.Points {
			q.Children[q.Bounds.WhichQuadrant(p)].Insert(p)
		}
		q.Points = nil
	}
	// Insert the point into the correct child
	q.Children[q.Bounds.WhichQuadrant(point)].Insert(point)
	q.TotalPoints++
}

// Removes the first occurance point from the quadtree.
// Will update the trees to the quadtree that initially called the function.
// @Returns the quadtree that the function starts from and the point that was removed.
func (q *Quadtree) RemovePoint(point Point) (*Quadtree, *Point) {
	if !q.Bounds.Contains(point) {
		return nil, nil
	}

	if q.IsLeaf {
		for i, p := range q.Points {
			if p == point {
				q.Points = append(q.Points[:i], q.Points[i+1:]...)
				q.TotalPoints--
				return q, &p
			}
		}
		return nil, nil
	}

	// Recursively search for the point
	q1, p := q.Children[q.Bounds.WhichQuadrant(point)].RemovePoint(point)
	if p == nil { // No child with match conditions was found
		return q1, p
	}

	// Found the child, now fix the tree
	q.TotalPoints--
	if q.TotalPoints < MaxPoints && q.Bounds.Width() > MinArea {
		// Merge the children
		points := make([]Point, 0)
		for _, child := range q.Children {
			if child != nil {
				points = append(points, child.Points...)
			}
		}
		q.Points = points
		q.IsLeaf = true
		for i := range q.Children {
			q.Children[i] = nil
		}

		q1 = q // Make it so that if the subtree was removed, the returned tree is the parent tree
	}

	return q1, p
}

/**
 * Moves a point from one quadtree to another
 * More accurately, takes in an old point, removes it, and inserts a new point at the new location
 * @Returns the quadtree the new point was inserted at? Might need to change
 * If the old point doesn't exist, won't insert the new point and returns null
 */
func (q *Quadtree) MovePoint(oldPoint Point, translation Point) *Quadtree {
	q = q.findMinimumBoundingTree([]Point{oldPoint, Translate(oldPoint, translation.X, translation.Y)})
	if q == nil {
		return nil
	}

	q1, _ := q.RemovePoint(oldPoint)
	if q1 == nil {
		return nil
	}
	oldPoint.Translate(translation.X, translation.Y)
	q.Insert(oldPoint)
	return q
}

/**
 * Finds the quadtree of the box that contains all the points in the array
 * @Returns nil if array is empty or points aren't all in the boundary
 */
func (q *Quadtree) findMinimumBoundingTree(points []Point) *Quadtree {
	bounds := NewBoundsWithPointArray(points)
	if bounds == nil {
		return nil
	}

	if !q.Bounds.Contains(bounds.BottomLeft) || !q.Bounds.Contains(bounds.TopRight) {
		return nil
	} else {
		if q.IsLeaf {
			return q
		}

		if q.Children[q.Bounds.WhichQuadrant(bounds.BottomLeft)].Bounds.Contains(bounds.TopRight) {
			return q.Children[q.Bounds.WhichQuadrant(bounds.BottomLeft)].findMinimumBoundingTree(points)
		}

		return q
	}
}

func (q *Quadtree) QueryRange(rangeBounds Bounds) []Point {
	// implement
	return nil
}

/**
 * Finds the quadtree the point should be in
 * @Returns nil if the point is not in the quadtree
 */
func (q *Quadtree) QueryPointQuadrant(point Point) *Quadtree {
	if !q.Bounds.Contains(point) {
		return nil
	}

	if q.IsLeaf {
		return q
	}

	// Recursively search for the point
	return q.Children[q.Bounds.WhichQuadrant(point)].QueryPointQuadrant(point)
}

/*
*
  - Finds all the points in the quadtree that are within a certain distance of the point

@Returns a QueuePoints with all the points that are within the distance of the point
*/
func (q *Quadtree) QueryNearby(point Point, dist int) *QueuePoints {
	queue := NewQueue()
	q.queryNearby(NewCircle(point, dist), queue)
	return queue
}
func (q *Quadtree) queryNearby(circle Circle, queue *QueuePoints) {
	if q.IsLeaf {
		for _, p := range q.Points {
			if circle.Contains(p) {
				queue.Enqueue(&p)
			}
		}
		return
	}
	for _, child := range q.Children {
		if child.Bounds.IntersectsCircle(circle) {
			child.queryNearby(circle, queue)
		}
	}
}

/*
Q1 (Bounds, isLeaf, TotalPoints): Points if is leaf
|
Q2 ...
|-- Q1 ...
|-- Q2 ...
|-- Q3 ...
|-- |-- Q1 ...
|-- |-- Q2 ...
|-- |-- Q3 ...
|-- |-- Q4 ...
|-- Q4 ...
|
Q3 ...
|
Q4 ...
*/
func (q *Quadtree) String() string {
	return q.stringWithIndent(0)
}

func (q *Quadtree) stringWithIndent(indent int) string {
	s := ""
	if q.IsLeaf {
		s += fmt.Sprintf("%sQ (%v, %v, %v): %v\n",
			strings.Repeat("|-- ", indent), q.Bounds, q.IsLeaf, q.TotalPoints, q.Points)
		return s
	}
	if indent == 0 {
		s += fmt.Sprintf("Q (%v, %v, %v)\n", q.Bounds, q.IsLeaf, q.TotalPoints)
		indent++
	}
	for i, child := range q.Children {
		s += fmt.Sprintf("%sQ%d (%v, %v, %v): ",
			strings.Repeat("|-- ", indent), i+1, child.Bounds, child.IsLeaf, child.TotalPoints)
		if child.IsLeaf {
			s += fmt.Sprintf("%v\n", child.Points)
		} else {
			s += fmt.Sprintf("\n%s", child.stringWithIndent(indent+1))
		}
	}

	return s
}
