package bst

import (
	"fmt"
	"testing"
)

func TestBST(t *testing.T) {
	bst := newBST()
	numbers := []int{5, 3, 6, 8, 4, 2}
	for _, val := range numbers {
		bst.Add(val)
	}
	fmt.Println("中序遍历:")
	fmt.Println(bst)

	fmt.Println("递归前序遍历")
	bst.PreOrder()
	fmt.Println()

	fmt.Println("非递归前序遍历")
	bst.PreOrderNR()

	fmt.Println()
	fmt.Println("广度优先遍历二分搜索树")
	bst.LevelOrder()
}
