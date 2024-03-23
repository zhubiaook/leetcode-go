package stack

import (
	"errors"
)

// Stacker is an interface for a stack
type Stacker interface {
	// Push adds an item to the top of the stack
	Push(item any)
	// Pop removes the top item from the stack
	// Returns an error if the stack is empty
	Pop() (any, error)
	// Peek returns the top item from the stack
	// Returns an error if the stack is empty
	Peek() (any, error)
	// Size returns the number of items in the stack
	Size() int
	// IsEmpty returns true if the stack is empty
	IsEmpty() bool
}

// New returns a new stack
func New() Stacker {
	s := make(stack, 0)
	return &s
}

// stack is a stack implementation using a slice
type stack []any

var _ Stacker = (*stack)(nil)

// Push adds an item to the top of the stack
func (s *stack) Push(item any) {
	*s = append(*s, item)
}

// Pop removes the top item from the stack
// Returns an error if the stack is empty
func (s *stack) Pop() (any, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	index := s.Size() - 1
	item := (*s)[index]
	*s = (*s)[:index]

	return item, nil
}

// Peek returns the top item from the stack
// Returns an error if the stack is empty
func (s *stack) Peek() (any, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	index := s.Size() - 1
	item := (*s)[index]

	return item, nil
}

// Size returns the number of items in the stack
func (s *stack) Size() int {
	return len(*s)
}

// IsEmpty returns true if the stack is empty, otherwise false
func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}
