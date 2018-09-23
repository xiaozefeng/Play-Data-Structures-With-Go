package linkedlist

import (
	"fmt"

	"github.com/xiaozefeng/Play-Data-Structures-With-Go/stack"
)

type LinkListStack struct {
	l *LinkedList
}

func (ls *LinkListStack) String() string {
	return fmt.Sprintln(ls.l)
}

func NewLinkListStack() stack.Stack {
	return &LinkListStack{
		NewLinkList(),
	}
}
func (ls *LinkListStack) Size() int {
	return ls.Size()
}

func (ls *LinkListStack) IsEmpty() bool {
	return ls.IsEmpty()
}

func (ls *LinkListStack) Push(e int) {
	ls.l.AddFirst(e)
}

func (ls *LinkListStack) Pop() int {
	return ls.l.RemoveFirst()
}

func (ls *LinkListStack) Peek() int {
	return ls.l.getFirst()
}
