package bst

import (
	"fmt"
	"math/rand"
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

	fmt.Printf("删除最小元素%d后树:%s\n", bst.RemoveMin(), bst)

	fmt.Printf("删除最大元素%d后树:%s\n", bst.RemoveMax(), bst)
}

func TestBST_RemoveMin(t *testing.T) {
	bst := newBST()
	n := 10000
	for i := 0; i < n; i++ {
		bst.Add(rand.Intn(n))
	}
	count := 0
	bst.traverseFunc(bst.root, func(node *node) {
		count++
		fmt.Printf("%d,", node.val)
	})
	fmt.Println()
	fmt.Println("count:", count)

	minSlice := make([]int, 0)
	for !bst.IsEmpty() {
		minSlice = append(minSlice, bst.RemoveMin())
	}
	for i := 1; i < len(minSlice); i++ {
		if minSlice[i-1] > minSlice[i] {
			panic("error")
		}
	}
}

func TestBST_RemoveMax(t *testing.T) {
	bst := newBST()
	n := 10000
	for i := 0; i < n; i++ {
		bst.Add(rand.Intn(n))
	}

	count := 0
	bst.traverseFunc(bst.root, func(node *node) {
		count++
		fmt.Printf("%d,", node.val)
	})
	fmt.Println()
	fmt.Println("count:", count)

	s := make([]int, 0)
	for !bst.IsEmpty() {
		s = append(s, bst.RemoveMax())
	}
	fmt.Println(s)
	for i := 1; i < len(s); i++ {
		if s[i-1] < s[i] {
			panic("error")
		}
	}
}
