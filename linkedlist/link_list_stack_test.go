package linkedlist

import (
	"fmt"
	"testing"
)

func TestLinkListStack(t *testing.T) {
	ls := NewLinkListStack()
	for i := 0; i < 5; i++ {
		ls.Push(i)
		fmt.Println(ls)
	}
	ls.Pop()
	ls.Pop()
	fmt.Println(ls)

}

func BenchmarkLinkedListStack(b *testing.B) {
	ls := NewLinkListStack()
	for i := 0; i < b.N; i++ {
		ls.Push(i)
	}
	for i := 0; i < b.N; i++ {
		ls.Pop()
	}
}
