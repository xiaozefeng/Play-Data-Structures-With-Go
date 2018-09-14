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
	fmt.Println(s)
	for i := 0; i < 5; i++ {
		s.Pop()
	}
	fmt.Println(s)
	s.Pop()
	fmt.Println(s)
	s.Pop()
}
