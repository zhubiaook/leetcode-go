package stack

// Stack is an interface for a stack data structure.
type Stack interface {
	// Push pushes a value onto the stack.
	Push(any)
	// Pop pops a value from the stack.
	Pop() (any, bool)
	// Top returns the top value of the stack.
	Top() (any, bool)
	// Size returns the number of elements in the stack.
	Size() int
	// IsEmpty returns true if the stack is empty.
	IsEmpty() bool
}
