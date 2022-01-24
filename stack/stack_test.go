package stack_test

import (
	"calculator/stack"
	"testing"
)

func TestStackEmpty(t *testing.T) {
	var s stack.Stack
	if !s.IsEmpty() {
		t.Errorf("Stack is not empty. Size = %d.", s.Size())
	}
}

func TestStackPush(t *testing.T) {
	var s stack.Stack

	s.Push(75)

	if s.IsEmpty() {
		t.Errorf("Stack is empty.")
	}
}

func TestStackSize(t *testing.T) {
	var s stack.Stack
	s.Push(1)
	s.Push(2)
	s.Push(3)

	want := 3

	if got := s.Size(); got != want {
		t.Errorf("want %d, got %d.", want, got)
	}
}

func TestStackPeekZeroElements(t *testing.T) {
	var s stack.Stack
	_, got := s.Peek()
	want := false
	if got {
		t.Errorf("want %t, got %t.", want, got)
	}
}

func TestStackPeek(t *testing.T) {
	var s stack.Stack
	var want float64 = 75

	s.Push(75)

	if got, _ := s.Peek(); got != want {
		t.Errorf("want %f, got %f.", want, got)
	}
}

func TestStackPeekiZeroElements(t *testing.T) {
	var s stack.Stack
	var want bool = false

	_, got := s.Peeki(0)

	if got != want {
		t.Errorf("want %t, got %t.", want, got)
	}
}

func TestStackPeekiMoreThanSize(t *testing.T) {
	var s stack.Stack
	s.Push(1)
	s.Push(2)
	s.Push(3)

	want := false
	_, got := s.Peeki(3)
	if want != got {
		t.Errorf("want %t, got %t.", want, got)
	}
}

func TestStackPeekiThreeElements(t *testing.T) {
	var s stack.Stack
	s.Push(1)
	s.Push(2)
	s.Push(3)

	var want float64 = 3
	got, ok := s.Peeki(0)
	if got != want {
		t.Errorf("want %f, got %f.", want, got)
	}
	if !ok {
		t.Errorf("want %t, got %t.", true, ok)
	}

	want = 2
	got, ok = s.Peeki(1)
	if got != want {
		t.Errorf("want %f, got %f.", want, got)
	}
	if !ok {
		t.Errorf("want %t, got %t.", true, ok)
	}

	want = 1
	got, ok = s.Peeki(2)
	if got != want {
		t.Errorf("want %f, got %f.", want, got)
	}
	if !ok {
		t.Errorf("want %t, got %t.", true, ok)
	}
}

func TestStackPopEmpty(t *testing.T) {
	var s stack.Stack
	want := false

	if _, got := s.Pop(); got != want {
		t.Errorf("want %t, got %t.", want, got)
	}
}

func TestStackPop(t *testing.T) {
	var s stack.Stack
	var want float64 = 75

	s.Push(75)
	got, ok := s.Pop()

	if !s.IsEmpty() {
		t.Errorf("Stack is not empty after pop. Size = %d.", s.Size())
	}

	if !ok {
		t.Errorf("Pop did not return ok.")
	}

	if got != want {
		t.Errorf("want %f, got %f.", want, got)
	}
}
