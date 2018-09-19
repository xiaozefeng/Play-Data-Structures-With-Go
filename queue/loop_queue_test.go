package queue

import (
	"fmt"
	"testing"
)

func TestNewLoopQueue(t *testing.T) {
	q := NewLoopQueue()
	for i := 0; i < 10; i++ {
		q.Enqueue(i)
		fmt.Println(q)
		if i%3 == 2 {
			q.Dequeue()
			fmt.Println(q)
		}
	}
}

func BenchmarkLoopQueue(b *testing.B) {
	q := NewLoopQueue()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
	for i := 0; i < b.N; i++ {
		q.Dequeue()
	}
}

func BenchmarkArrayQueue(b *testing.B) {
	q := NewArrayQueue()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
	for i := 0; i < b.N; i++ {
		q.Dequeue()
	}
}
