package set

import (
	"fmt"
	"strings"
)

type Set interface {
	Add(e string)
	Remove(e string)
	Size() int
	IsEmpty() bool
	Contains(e string) bool
}

func NewBSTSet() Set {
	return &BSTSet{bst: newBst()}
}

type BSTSet struct {
	bst *bst
}

func (b *BSTSet) Add(e string) {
	b.bst.add(e)
}

func (b *BSTSet) Remove(e string) {
	b.bst.remove(e)
}

func (b *BSTSet) Size() int {
	return b.bst.getSize()
}

func (b *BSTSet) IsEmpty() bool {
	return b.bst.isEmpty()
}

func (b *BSTSet) Contains(e string) bool {
	return b.bst.contains(e)
}

type node struct {
	val         string
	left, right *node
}

func newNode(e string) *node {
	return &node{val: e}
}

type bst struct {
	root *node
	size int
}

func newBst() *bst {
	return &bst{
		root: nil,
		size: 0,
	}
}

func (b *bst) getSize() int {
	return b.size
}

func (b *bst) isEmpty() bool {
	return b.size == 0
}

func (b *bst) traverseFunc(node *node, f func(*node)) {
	if node == nil {
		return
	}
	b.traverseFunc(node.left, f)
	f(node)
	b.traverseFunc(node.right, f)
}

func (b *bst) contains(e string) bool {
	result := false
	b.traverseFunc(b.root, func(node *node) {
		if e == node.val {
			result = true
			return
		}
	})
	return result
}

func (b *bst) add(e string) {
	b.root = b.addEle(b.root, e)
}

// 向node 为根的树中添加元素，并且返回添加后的根
func (b *bst) addEle(node *node, e string) *node {
	if node == nil {
		b.size++
		return newNode(e)
	}
	compare := strings.Compare(e, node.val)
	if compare == -1 {
		node.left = b.addEle(node.left, e)
	} else if compare == 1 {
		node.right = b.addEle(node.right, e)
	}
	return node
}

func (b *bst) minimum() string {
	if b.size == 0 {
		panic("bst is empty")
	}
	return b.min(b.root).val
}

func (b *bst) min(node *node) *node {
	if node.left == nil {
		return node
	}
	return b.min(node.left)
}

func (b *bst) maximum() string {
	if b.size == 0 {
		panic("bst is empty.")
	}
	return b.max(b.root).val
}

func (b *bst) max(node *node) *node {
	if node.right == nil {
		return node
	}
	return b.max(node.right)
}

func (b *bst) removeMinimum() *node {
	if b.size == 0 {
		panic("bst is empty")
	}
	ret := b.min(b.root)
	b.root = b.removeMin(b.root)
	return ret
}

func (b *bst) removeMin(node *node) *node {
	if node.left == nil {
		right := node.right
		node.right = nil
		b.size--
		return right
	}
	node.left = b.removeMin(node.right)
	return node
}

func (b *bst) removeMaximum() *node {
	if b.size == 0 {
		panic("bst is empty")
	}
	max := b.max(b.root)
	b.root = b.removeMax(b.root)
	return max
}

func (b *bst) removeMax(node *node) *node {
	if node.right == nil {
		left := node.left
		node.left = nil
		b.size--
		return left
	}
	node.right = b.removeMax(node.right)
	return node
}

func (b *bst) remove(e string) {

}

func (b *bst) removeElement(node *node, e string) *node {
	if node == nil {
		return nil
	}
	compare := strings.Compare(e, node.val)
	if compare == -1 {
		node.left = b.removeElement(node.left, e)
		return node
	} else if compare == 1 {
		node.right = b.removeElement(node.right, e)
		return node
	} else {
		// 待删除元素左子树为空的情况
		if node.left == nil {
			right := node.right
			node.right = nil
			b.size--
			return right
		}
		// 待删除元素右子树为空的情况
		if node.right == nil {
			left := node.left
			node.left = nil
			b.size--
			return left
		}

		// 待删除元素左右子树都不为空的情况
		// 找到当前元素的后继元素
		successor := b.min(node.right)
		successor.left = node.left
		successor.right = b.removeMin(node.right)
		node.left = nil
		node.right = nil
		return successor
	}
}

func (b *bst) String() string {
	result := ""
	b.traverseFunc(b.root, func(node *node) {
		result += fmt.Sprintf("%s, ", node.val)
	})
	return result
}
