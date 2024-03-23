package stack

import (
	"testing"
)

func TestNew(t *testing.T) {
	s := New()
	if s == nil {
		t.Errorf("New() returned a nil stack")
	}
	if !s.IsEmpty() {
		t.Errorf("New() returned a non-empty stack")
	}
}

func TestPush(t *testing.T) {
	s := New()
	s.Push(42)
	if s.Size() != 1 {
		t.Errorf("Size() should be 1 after Push()")
	}
	if s.IsEmpty() {
		t.Errorf("Stack should not be empty after Push()")
	}
}

func TestPop(t *testing.T) {
	s := New()
	item, err := s.Pop()
	if err == nil {
		t.Errorf("Pop() should return an error on an empty stack")
	}

	s.Push("foo")
	item, err = s.Pop()
	if err != nil {
		t.Errorf("Pop() returned an error: %v", err)
	}
	if item != "foo" {
		t.Errorf("Pop() returned incorrect item: %v", item)
	}
	if !s.IsEmpty() {
		t.Errorf("Stack should be empty after Pop()")
	}
}

func TestPeek(t *testing.T) {
	s := New()
	item, err := s.Peek()
	if err == nil {
		t.Errorf("Peek() should return an error on an empty stack")
	}

	s.Push(10)
	item, err = s.Peek()
	if err != nil {
		t.Errorf("Peek() returned an error: %v", err)
	}
	if item != 10 {
		t.Errorf("Peek() returned incorrect item: %v", item)
	}
	if s.Size() != 1 {
		t.Errorf("Size() should be 1 after Peek()")
	}
}

func TestSize(t *testing.T) {
	s := New()
	if s.Size() != 0 {
		t.Errorf("Size() should be 0 for an empty stack")
	}

	s.Push("foo")
	s.Push("bar")
	if s.Size() != 2 {
		t.Errorf("Size() should be 2 after pushing 2 items")
	}
}

func TestIsEmpty(t *testing.T) {
	s := New()
	if !s.IsEmpty() {
		t.Errorf("IsEmpty() should return true for an empty stack")
	}

	s.Push("foo")
	if s.IsEmpty() {
		t.Errorf("IsEmpty() should return false for a non-empty stack")
	}
}
