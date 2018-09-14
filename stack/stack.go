package stack

import (
	"fmt"

	"github.com/xiaozefeng/Play-Data-Structures-With-Go/array"
)

// Stack data struct
type Stack interface {
	Size() int
	IsEmpty() bool
	Push(e int)
	Pop() int
	Peek() int
}

// ArrayStack data struct
type ArrayStack struct {
	arr *array.Array
}

// Size commonet
func (st *ArrayStack) Size() int {
	return st.arr.Size()
}

// IsEmpty comment
func (st *ArrayStack) IsEmpty() bool {
	return st.arr.IsEmpty()
}

// Push comment
func (st *ArrayStack) Push(e int) {
	st.arr.AddLast(e)
}

// Pop comment
func (st *ArrayStack) Pop() int {
	ret, _ := st.arr.RemoveLast()
	return ret
}

// Peek comment
func (st *ArrayStack) Peek() int {
	ret, _ := st.arr.GetLast()
	return ret

}

// NewStack create a satck
func NewStack() Stack {
	return &ArrayStack{
		arr: array.NewArray(),
	}
}

func (st *ArrayStack) String() string {
	v := ""
	v += "[Stack: "
	for i := st.arr.Size() - 1; i >= 0; i-- {
		r, _ := st.arr.Get(i)
		v += fmt.Sprintf("%v ", r)
	}
	v += "]"
	return v
}
