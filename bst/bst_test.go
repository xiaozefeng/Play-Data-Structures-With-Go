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

	fmt.Println()
	fmt.Println("二分搜索树的最大值", bst.MaxValueNR())
	fmt.Println("二分搜索树的最大值, 递归算法", bst.Max())

	fmt.Println("二分搜索树的最小值", bst.MinValueNR())
	fmt.Println("二分搜索树的最小值, 递归算法", bst.Minimum())

	bst.RemoveMin()
	fmt.Println("删除最小元素后", bst)

	bst.RemoveMax()
	fmt.Println("删除最大元素后", bst)

}
