package array

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	arr := NewArray()
	err := arr.Add(0, 1)
	if err != nil {
		panic(err)
	}
	arr.Add(1, 1)
	arr.Add(0, 2)
	arr.Add(2, 3)
	fmt.Printf("%s\n", arr)
	// 2131
	e, er := arr.Remove(0)
	if er != nil {
		panic(er)
	}
	fmt.Printf("被删除的元素:%d\n", e)
}
