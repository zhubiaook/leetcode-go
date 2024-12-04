package queue

// sliceQueue implements the Queue interface
type sliceQueue []any

var _ Queue = new(sliceQueue)

// NewSliceQueue creates a new slice-based queue.
func NewSliceQueue() Queue {
	return &sliceQueue{}
}

// Enqueue enqueues a value to the queue.
func (q *sliceQueue) Enqueue(v any) {
	(*q) = append(*q, v)
}

// Dequeue dequeues a value from the queue.
func (q *sliceQueue) Dequeue() (any, bool) {
	if q.IsEmpty() {
		return nil, false
	}
	e := (*q)[0]
	(*q) = (*q)[1:]
	return e, true
}

// Front returns the front value of the queue.
func (q *sliceQueue) Front() (any, bool) {
	if q.IsEmpty() {
		return nil, false
	}
	return (*q)[0], true
}

// Size returns the number of elements in the queue.
func (q *sliceQueue) Size() int {
	return len(*q)
}

// IsEmpty returns true if the queue is empty.
func (q *sliceQueue) IsEmpty() bool {
	return len(*q) == 0
}
