package queue

import (
	"testing"
)

func TestSliceQueue(t *testing.T) {
	// Create a new queue instance
	q := NewSliceQueue()

	// Test: Queue should be empty initially
	if !q.IsEmpty() {
		t.Errorf("Expected queue to be empty, but it is not")
	}

	// Test: Queue size should be 0
	if q.Size() != 0 {
		t.Errorf("Expected queue size to be 0, got %d", q.Size())
	}

	// Enqueue values and check size
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if q.Size() != 3 {
		t.Errorf("Expected queue size to be 3, got %d", q.Size())
	}

	// Test: Front should return the first enqueued value
	front, ok := q.Front()
	if !ok || front != 1 {
		t.Errorf("Expected front to be 1, got %v", front)
	}

	// Dequeue values and check correctness
	deq1, ok := q.Dequeue()
	if !ok || deq1 != 1 {
		t.Errorf("Expected dequeued value to be 1, got %v", deq1)
	}

	deq2, ok := q.Dequeue()
	if !ok || deq2 != 2 {
		t.Errorf("Expected dequeued value to be 2, got %v", deq2)
	}

	// Check if the queue is not empty yet
	if q.IsEmpty() {
		t.Errorf("Expected queue to be non-empty")
	}

	// Dequeue the last value
	_, _ = q.Dequeue()

	// Check if the queue is empty
	if !q.IsEmpty() {
		t.Errorf("Expected queue to be empty")
	}

	// Dequeue from an empty queue
	_, ok = q.Dequeue()
	if ok {
		t.Errorf("Expected Dequeue on empty queue to return false")
	}

	// Front from an empty queue
	_, ok = q.Front()
	if ok {
		t.Errorf("Expected Front on empty queue to return false")
	}
}
