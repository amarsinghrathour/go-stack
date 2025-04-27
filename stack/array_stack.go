package stack

// ArrayStack is a stack implementation using a slice (array-like) in Go
type ArrayStack struct {
	items []interface{}
}

// NewArrayStack creates a new ArrayStack
func NewArrayStack() *ArrayStack {
	return &ArrayStack{}
}

// Push adds an element to the stack
func (s *ArrayStack) Push(value interface{}) {
	s.items = append(s.items, value)
}

// Pop removes and returns the top element of the stack
func (s *ArrayStack) Pop() (interface{}, bool) {
	if len(s.items) == 0 {
		return nil, false
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true
}

// Peek returns the top element without removing it
func (s *ArrayStack) Peek() (interface{}, bool) {
	if len(s.items) == 0 {
		return nil, false
	}
	return s.items[len(s.items)-1], true
}

// IsEmpty checks if the stack is empty
func (s *ArrayStack) IsEmpty() bool {
	return len(s.items) == 0
}
