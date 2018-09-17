package queue

import (
	"fmt"
)

type Queue interface {
	Size() int
	IsEmpty() bool
	Enqueue(e int)
	Dequeue() int
	GetFront() int
}

func NewArrayQueue() Queue {
	return &ArrayQueue{
		data: make([]int, 0),
		size: 0,
	}
}

type ArrayQueue struct {
	data []int
	size int
}

func (q *ArrayQueue) Size() int {
	return q.size
}

func (q *ArrayQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *ArrayQueue) Enqueue(e int) {
	q.data = append(q.data, e)
	q.size++
}

func (q *ArrayQueue) Dequeue() int {
	ret := q.data[1]
	q.data = q.data[1:]
	q.size--
	return ret
}

func (q *ArrayQueue) GetFront() int {
	return q.data[1]
}

func (q ArrayQueue) String() string {
	s := "Queue: front["
	for i := 0; i < q.size; i++ {
		s += fmt.Sprintf("%v, ", q.data[i])
	}
	s += "]tail"
	return s
}
