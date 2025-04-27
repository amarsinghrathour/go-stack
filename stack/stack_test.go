package stack

import "testing"

func TestArrayStack(t *testing.T) {
	s := NewArrayStack()
	s.Push(1)
	s.Push(2)
	if v, _ := s.Pop(); v != 2 {
		t.Errorf("Expected 2, got %v", v)
	}
}

func TestLinkedStack(t *testing.T) {
	s := NewLinkedStack()
	s.Push("a")
	s.Push("b")
	if v, _ := s.Pop(); v != "b" {
		t.Errorf("Expected 'b', got %v", v)
	}
}

func TestConcurrentStack(t *testing.T) {
	s := NewConcurrentStack()
	s.Push(10)
	s.Push(20)
	if v, _ := s.Pop(); v != 20 {
		t.Errorf("Expected 20, got %v", v)
	}
}
