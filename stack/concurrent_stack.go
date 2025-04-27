package stack

import "sync"

// ConcurrentStack is a thread-safe stack using sync.Mutex
type ConcurrentStack struct {
	items []interface{}
	mu    sync.Mutex
}

// NewConcurrentStack creates a new ConcurrentStack
func NewConcurrentStack() *ConcurrentStack {
	return &ConcurrentStack{}
}

// Push adds an element to the stack in a thread-safe manner
func (s *ConcurrentStack) Push(value interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = append(s.items, value)
}

// Pop removes and returns the top element of the stack in a thread-safe manner
func (s *ConcurrentStack) Pop() (interface{}, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.items) == 0 {
		return nil, false
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true
}

// Peek returns the top element without removing it
func (s *ConcurrentStack) Peek() (interface{}, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.items) == 0 {
		return nil, false
	}
	return s.items[len(s.items)-1], true
}

// IsEmpty checks if the stack is empty
func (s *ConcurrentStack) IsEmpty() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.items) == 0
}
