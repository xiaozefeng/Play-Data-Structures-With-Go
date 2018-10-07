package bst

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	bst := newBST()
	numbers := []int{5, 3, 6, 8, 4, 2}
	for _, val := range numbers {
		bst.Add(val)
	}
	fmt.Println("中序遍历:")
	fmt.Println(bst)
}
