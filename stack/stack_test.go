package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack()
	for i := 0; i < 6; i++ {
		s.Push(i)
	}
	expected := 6
	if s.Size() != expected {
		t.Errorf("expected %d , actual got %d", expected, s.Size())
	}
	fmt.Println(s)
	for i := 0; i < 5; i++ {
		s.Pop()
	}
	if s.Peek() != 0 {
		t.Errorf("expected 0, actual got %d", s.Peek())
	}
	fmt.Println(s)
	s.Pop()
	fmt.Println(s)
}
