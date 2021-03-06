package bst

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestBST(t *testing.T) {
	bst := New()
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
	bst := New()
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
	bst := New()
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

func TestBST_Remove(t *testing.T) {
	bst := New()
	for i := 0; i < 10; i++ {
		bst.Add(i)
	}
	fmt.Println(bst)

	bst.Remove(5)
	fmt.Println("bst remove element: ", bst)
}

func TestBST_Floor(t *testing.T) {
	bst := New()
	bst.Add(5)
	bst.Add(6)
	bst.Add(10)
	bst.Add(1)
	fmt.Println("Floor:", bst.Floor(7))
}

func TestBST_Rank(t *testing.T) {
	bst := New()
	bst.Add(41)
	bst.Add(22)
	bst.Add(15)
	bst.Add(13)
	bst.Add(33)
	bst.Add(37)
	bst.Add(58)
	bst.Add(50)
	bst.Add(42)
	bst.Add(53)
	bst.Add(63)
	fmt.Println("58的排名是", bst.Rank(58))
	fmt.Println("排名10的元素是", bst.Select(10))
}
