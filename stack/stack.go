// Package stack provides various stack implementations.
package stack

// Stack is the interface that all stack implementations must fulfill.
type Stack interface {
	// Push adds an item to the top of the stack.
	Push(value interface{})

	// Pop removes and returns the item from the top of the stack.
	// It returns an error if the stack is empty.
	Pop() (interface{}, bool)

	// Peek returns the item at the top of the stack without removing it.
	// It returns an error if the stack is empty.
	Peek() (interface{}, bool)

	// IsEmpty checks if the stack is empty.
	IsEmpty() bool
}
