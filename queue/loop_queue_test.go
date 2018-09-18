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
