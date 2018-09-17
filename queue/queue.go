package queue

type Queue interface {
	Size() int
	IsEmpty() bool
	Enqueue(e int)
	Dequeue() int
	GetFront() int
}
