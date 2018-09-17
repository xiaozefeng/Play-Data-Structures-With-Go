package queue

import "fmt"

type LoopQueue struct {
	data              []int
	front, tail, szie int
}

func NewLoopQueue() Queue {
	return NewLoopQueueWithCapacity(10)
}

// we are consciously wasting one  space
// because this is a loop queue
// this queue has front and tail node, the point to the head and tail respecitively
func NewLoopQueueWithCapacity(capacity int) Queue {
	return &LoopQueue{
		data:  make([]int, capacity+1, capacity+1),
		front: 0,
		tail:  0,
		szie:  0,
	}
}

func (q *LoopQueue) Size() int {
	return q.szie
}

// while the front == tail , the queue is empty
func (q *LoopQueue) IsEmpty() bool {
	return q.front == q.tail
}

// while the
func (q *LoopQueue) Enqueue(e int) {
	// judging queue capacity
	if (q.tail+1)%len(q.data) == q.front {
		q.resize(q.getCapacity() + 1)
	}
	q.data[q.tail] = e
	// because this queue is loop queue, use (index +1)% data.length to get correct index
	q.tail = (q.tail + 1) % len(q.data)
	q.szie++
}

func (q *LoopQueue) Dequeue() int {
	if q.IsEmpty() {
		panic("Cannot dequeue from an empty queue.")
	}
	ret := q.data[q.front]
	q.front = (q.front + 1) % len(q.data)
	q.szie--
	// reduce capacity by half when only a quarter of the space is utilized
	if q.Size() == q.getCapacity()/4 && q.getCapacity()/2 != 0 {
		q.resize(len(q.data) / 2)
	}
	return ret
}

func (q *LoopQueue) GetFront() int {
	if q.IsEmpty() {
		panic("Cannot dequeue from an empty queue.")
	}
	return q.data[q.front]
}

func (q *LoopQueue) getCapacity() int {
	return len(q.data) - 1
}

// the private method
// resize LoopQueue's data capacity
func (q *LoopQueue) resize(newCapacity int) {
	newData := make([]int, newCapacity+1, newCapacity+1)
	for i := 0; i < q.Size(); i++ {
		// put the `front`+ i  position of origin array into the new array
		lenght := len(q.data)
		newData[i] = q.data[(q.front+i)%lenght]
	}
	q.data = newData
	q.front = 0
	q.tail = q.Size()
}

func (q *LoopQueue) String() string {
	s := "Queue: front["
	for i := q.front; i != q.tail; i = (i + 1) % len(q.data) {
		s += fmt.Sprintf("%v, ", q.data[i])
	}
	s += "]tail"
	return s
}
