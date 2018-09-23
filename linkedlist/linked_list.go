package linkedlist

import (
	"fmt"
	"strconv"
)

type node struct {
	value int
	next  *node
}

func (n node) String() string {
	return strconv.Itoa(n.value)
}

// create a node and init value
func newNode(e int) *node {
	return newNodeWithValueAndNext(e, nil)
}

// create a node and init value and next node
func newNodeWithValueAndNext(e int, next *node) *node {
	return &node{
		value: e,
		next:  next,
	}
}

type LinkedList struct {
	size      int
	dummyHead *node
}

func NewLinkList() *LinkedList {
	return &LinkedList{
		size:      0,
		dummyHead: newNode(-1),
	}
}
func (l *LinkedList) String() string {
	str := ""
	for cur := l.dummyHead.next; cur != nil; cur = cur.next {
		str += fmt.Sprintf("%d ->", cur.value)
	}
	str += "nil"
	return str
}

func (l *LinkedList) Size() int {
	return l.size
}
func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}
func (l *LinkedList) Add(index int, e int) {
	// check params
	if index < 0 || index > l.size {
		panic("Add field . Illegal index")
	}
	// deal special situation, index == 0 , don't have previous node
	pre := l.dummyHead
	for i := 0; i < index; i++ {
		// move the node
		pre = pre.next
	}
	// node := newNode(e)
	// node.next = pre.next
	// pre.next = node
	pre.next = newNodeWithValueAndNext(e, pre.next)
	l.size++
}

func (l *LinkedList) AddFirst(e int) {
	l.Add(0, e)
}

func (l *LinkedList) AddLast(e int) {
	l.Add(l.size, e)
}

// practice
func (l *LinkedList) Get(index int) int {
	if index < 0 || index >= l.size {
		panic("Get failed ,Illegal index")
	}

	cur := l.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.value
}

func (l *LinkedList) getFirst() int {
	return l.Get(0)
}

func (l *LinkedList) GetLast() int {
	return l.Get(l.size - 1)
}

func (l *LinkedList) Set(index, e int) {
	if index < 0 || index >= l.size {
		panic("Set failed. Illegal index")
	}
	cur := l.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	cur.value = e
}

func (l *LinkedList) Contains(e int) bool {
	if l.IsEmpty() {
		return false
	}
	for cur := l.dummyHead.next; cur != nil; cur = cur.next {
		if e == cur.value {
			return true
		}
	}
	return false
}

func (l *LinkedList) Remove(index int) (e int) {
	if index < 0 || index >= l.size {
		panic("Remove failed. Illegal index")
	}
	pre := l.dummyHead
	for i := 0; i < index; i++ {
		pre = pre.next
	}
	delNode := pre.next
	pre.next = delNode.next
	// delNode.next = nil
	l.size--
	return delNode.value
}

func (l *LinkedList) RemoveFirst() (e int) {
	return l.Remove(0)
}

func (l *LinkedList) RemoveLast() (e int) {
	return l.Remove(l.size - 1)
}
