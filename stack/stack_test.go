package stack

import "testing"

func TestPushLength(t *testing.T) {
	stack := New()
	stack.Push(1)

	if stack.Length() != 1 {
		t.Errorf("Length must be 1, got %d", stack.Length())
	}
}

func TestPush(t *testing.T) {
	stack := New()
	stack.Push(2)

	if stack.Top() != 2 {
		t.Errorf("Top must be 2, got %d", stack.Top())
	}
}

func TestPop(t *testing.T) {
	stack := New()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	temp := stack.Pop()
	if temp != 3 {
		t.Errorf("Pop should return 3, got %d", temp)
	}
}

func TestAfterPop(t *testing.T) {
	stack := New()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	stack.Pop()

	if stack.Length() != 2 {
		t.Errorf("Length should be 2, got %d", stack.Length())
	}
}
