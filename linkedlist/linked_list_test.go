package linkedlist

import (
	"fmt"
	"testing"
)

func TestLinkList(t *testing.T) {
	l := NewLinkList()
	for i := 0; i < 5; i++ {
		l.AddFirst(i)
		fmt.Println(l)
	}
	l.Add(2, 666)
	fmt.Println(l)
	delValue := l.Remove(2)
	fmt.Printf("del value :%d\n", delValue)
	fmt.Println(l)

	l.RemoveFirst()
	fmt.Println(l)
	l.RemoveLast()
	fmt.Println(l)
}
