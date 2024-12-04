package stack_test

import (
	"testing"

	"github.com/zhubiaook/leetcode-go/go-data-structure/stack"
)

// TestSliceStack tests the slice-based stack.
func TestSliceStack(t *testing.T) {
	// Create a new stack instance
	s := stack.NewSliceStack()

	// Test: Stack should be empty initially
	if !s.IsEmpty() {
		t.Errorf("Expected stack to be empty, but it is not")
	}

	// Test: Stack size should be 0
	if s.Size() != 0 {
		t.Errorf("Expected stack size to be 0, got %d", s.Size())
	}

	// Push values and check size
	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.Size() != 3 {
		t.Errorf("Expected stack size to be 3, got %d", s.Size())
	}

	// Test: Top should return the last pushed value
	top, ok := s.Top()
	if !ok || top != 3 {
		t.Errorf("Expected top to be 3, got %v", top)
	}

	// Pop values and check correctness
	pop1, ok := s.Pop()
	if !ok || pop1 != 3 {
		t.Errorf("Expected popped value to be 3, got %v", pop1)
	}

	pop2, ok := s.Pop()
	if !ok || pop2 != 2 {
		t.Errorf("Expected popped value to be 2, got %v", pop2)
	}

	// Check if the stack is not empty yet
	if s.IsEmpty() {
		t.Errorf("Expected stack to be non-empty")
	}

	// Pop the last value
	_, _ = s.Pop()

	// Check if the stack is empty
	if !s.IsEmpty() {
		t.Errorf("Expected stack to be empty")
	}

	// Pop from an empty stack
	_, ok = s.Pop()
	if ok {
		t.Errorf("Expected Pop on empty stack to return false")
	}

	// Top from an empty stack
	_, ok = s.Top()
	if ok {
		t.Errorf("Expected Top on empty stack to return false")
	}
}
