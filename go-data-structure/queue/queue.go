package queue

// Queue is an interface for a queue data structure.
type Queue interface {
	// Enqueue enqueues a value to the queue.
	Enqueue(any)
	// Dequeue dequeues a value from the queue.
	Dequeue() (any, bool)
	// Front returns the front value of the queue.
	Front() (any, bool)
	// Size returns the number of elements in the queue.
	Size() int
	// IsEmpty returns true if the queue is empty.
	IsEmpty() bool
}
