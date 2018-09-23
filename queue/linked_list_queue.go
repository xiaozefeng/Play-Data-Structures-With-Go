package queue

import (
	"fmt"
	"strconv"
)

type node struct {
	value int
	next  *node
}

func newNode(e int) *node {
	return &node{
		value: e,
	}
}

func (n *node) String() string {
	return strconv.Itoa(n.value)
}

type LinkedListQueue struct {
	size       int
	head, tail *node
}

func NewLinkedListQueue() Queue {
	return &LinkedListQueue{
		size: 0,
		head: nil,
		tail: nil,
	}
}

func (q *LinkedListQueue) Size() int {
	return q.size
}

func (q *LinkedListQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *LinkedListQueue) Enqueue(e int) {
	if q.tail == nil {
		q.tail = newNode(e)
		q.head = q.tail
	} else {
		q.tail.next = newNode(e)
		q.tail = q.tail.next
	}
	q.size++
}

func (q *LinkedListQueue) Dequeue() int {
	if q.IsEmpty() {
		panic("Dequeue failed. queues is empty")
	}
	delNode := q.head
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	q.size--
	return delNode.value
}

func (q *LinkedListQueue) GetFront() int {
	if q.IsEmpty() {
		panic("Dequeue failed. queues is empty")
	}
	return q.head.value
}

func (q *LinkedListQueue) String() string {
	str := "Queue: front "
	for cur := q.head; cur != nil; cur = cur.next {
		str += fmt.Sprintf("%d ->", cur.value)
	}
	str += "nil tail"
	return str
}
