package stack

import "testing"

func TestPushLength(t *testing.T) {
	s := &Stack{}
	s.Push(1)

	if s.Length() != 1 {
		t.Errorf("Length must be 1, got %d", s.Length())
	}
}

func TestPush(t *testing.T) {
	s := &Stack{}
	s.Push(2)

	if s.Peek() != 2 {
		t.Errorf("Peek must be 2, got %d", s.Peek())
	}
}

func TestPop(t *testing.T) {
	s := &Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)

	temp := s.Pop()
	if temp != 3 {
		t.Errorf("Pop should return 3, got %d", temp)
	}
}

func TestAfterPop(t *testing.T) {
	s := &Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)

	s.Pop()

	if s.Length() != 2 {
		t.Errorf("Length should be 2, got %d", s.Length())
	}
}

func TestIsEmpty(t *testing.T) {
	s := &Stack{}
	s.Push(1)

	if s.IsEmpty() {
		t.Error("Stack should not be empty")
	}

	s.Pop()

	if !s.IsEmpty() {
		t.Error("Stack should be empty")
	}

}
