package bst

import (
	"fmt"
)

type node struct {
	val         int
	left, right *node
}

func newNode(val int) *node {
	return &node{
		val: val,
	}
}

type BST struct {
	root *node
	size int
}

func (b *BST) String() string {
	str := ""
	b.traverseFunc(b.root, func(n *node) {
		str += fmt.Sprintf("%d,", n.val)
	})
	return str
}

func newBST() *BST {
	return &BST{
		root: nil,
		size: 0,
	}
}

//Size
func (b *BST) Size() int {
	return b.size
}

//IsEmpty
func (b *BST) IsEmpty() bool {
	return b.size == 0
}

func (b *BST) Add(e int) {
	b.root = b.add(b.root, e)
}

// 向以nod{}e 为跟的二分搜索树中插入元素e , 递归算法
// 返回插入新节点后的二分搜索树的根
func (b *BST) add(n *node, e int) *node {
	if n == nil {
		return newNode(e)
	}
	if e < n.val {
		n.left = b.add(n.left, e)
	} else {
		n.right = b.add(n.right, e)
	}
	return n
}

func (b *BST) Contains(e int) bool {
	return b.contains(b.root, e)
}

func (b *BST) contains(n *node, e int) bool {
	if n == nil {
		return false
	}
	if n.val == e {
		return true
	}

	if e < n.val {
		return b.contains(n.left, e)
	} else {
		return b.contains(n.right, e)
	}

}

func (b *BST) PreOrder() {
	b.traverse(b.root)
}

func (b *BST) traverseFunc(n *node, f func(n *node)) {
	if n == nil {
		return
	}
	b.traverseFunc(n.left, f)
	f(n)
	b.traverseFunc(n.right, f)
}

func (b *BST) traverse(n *node) {
	if n == nil {
		return
	}
	fmt.Printf("%d, ", n.val)
	b.traverse(n.left)
	b.traverse(n.right)
}
