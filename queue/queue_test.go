package queue

import (
	"fmt"
	"testing"
)

func TestArrayQueue(t *testing.T) {
	q := NewArrayQueue()
	for i := 0; i < 10; i++ {
		q.Enqueue(i)
		fmt.Println(q)

		if i%3 == 2 {
			q.Dequeue()
			fmt.Println(q)
		}
	}
}
