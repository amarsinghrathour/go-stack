package stack

// Node represents a single item in the linked list
type Node struct {
	value interface{}
	next  *Node
}

// LinkedStack is a stack implemented with a linked list
type LinkedStack struct {
	top *Node
}

// NewLinkedStack creates a new LinkedStack
func NewLinkedStack() *LinkedStack {
	return &LinkedStack{}
}

// Push adds an element to the stack
func (s *LinkedStack) Push(value interface{}) {
	newNode := &Node{value: value, next: s.top}
	s.top = newNode
}

// Pop removes and returns the top element of the stack
func (s *LinkedStack) Pop() (interface{}, bool) {
	if s.top == nil {
		return nil, false
	}
	topValue := s.top.value
	s.top = s.top.next
	return topValue, true
}

// Peek returns the top element without removing it
func (s *LinkedStack) Peek() (interface{}, bool) {
	if s.top == nil {
		return nil, false
	}
	return s.top.value, true
}

// IsEmpty checks if the stack is empty
func (s *LinkedStack) IsEmpty() bool {
	return s.top == nil
}
