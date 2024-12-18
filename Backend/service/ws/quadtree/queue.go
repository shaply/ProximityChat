package quadtree

// This is the queue for the query nearby function

type node struct {
	point *Point
	next  *node
}

func NewNode(point *Point) *node {
	return &node{point: point}
}

type QueuePoints struct {
	head *node
	tail *node
	size int
}

func NewQueue() *QueuePoints {
	h := NewNode(nil)
	return &QueuePoints{
		head: h,
		tail: h,
		size: 0,
	}
}

func (q *QueuePoints) Enqueue(point *Point) {
	q.tail.next = NewNode(point)
	q.tail = q.tail.next
	q.size++
}

func (q *QueuePoints) Dequeue() *Point {
	if q.size == 0 {
		return nil
	}

	point := q.head.next.point
	q.head.next = q.head.next.next
	q.size--
	if q.size == 0 {
		q.tail = q.head
	}
	return point
}

type QueuePointsIterator struct {
	current *node
}

func (q *QueuePoints) Iterator() *QueuePointsIterator {
	return &QueuePointsIterator{current: q.head.next}
}

func (it *QueuePointsIterator) HasNext() bool {
	return it.current != nil
}

func (it *QueuePointsIterator) Next() *Point {
	if !it.HasNext() {
		return nil
	}
	point := it.current.point
	it.current = it.current.next
	return point
}

func (q *QueuePoints) Range() <-chan *Point {
	ch := make(chan *Point)
	go func() {
		for it := q.Iterator(); it.HasNext(); {
			ch <- it.Next()
		}
		close(ch)
	}()
	return ch
}
