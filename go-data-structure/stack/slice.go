package stack

// sliceStack implements the Stack interface
type sliceStack []any

var _ Stack = new(sliceStack)

// NewSliceStack creates a new slice-based stack.
func NewSliceStack() Stack {
	return &sliceStack{}
}

// Push pushes a value onto the stack.
func (s *sliceStack) Push(v any) {
	*s = append(*s, v)
}

// Pop pops a value from the stack.
func (s *sliceStack) Pop() (any, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	e := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return e, true
}

// Top returns the top value of the stack.
func (s *sliceStack) Top() (any, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	return (*s)[len(*s)-1], true
}

// Size returns the number of elements in the stack.
func (s *sliceStack) Size() int {
	return len(*s)
}

// IsEmpty returns true if the stack is empty.
func (s *sliceStack) IsEmpty() bool {
	return len(*s) == 0
}
