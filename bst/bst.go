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

type stack struct {
	data []*node
	size int
}

func (s *stack) getSize() int {
	return s.size
}

func newStack() *stack {
	return &stack{
		size: 0,
		data: make([]*node, 0),
	}
}

func (s *stack) isEmpty() bool {
	return s.size == 0
}

func (s *stack) push(node *node) {
	s.data = append(s.data, node)
	s.size++
}

func (s *stack) pop() *node {
	if s.isEmpty() {
		panic("pop failed. stack is empty")
	}
	stackLen := len(s.data)
	ret := s.data[stackLen-1]
	s.data = s.data[:stackLen-1]
	s.size--
	return ret
}

type queue struct {
	data []*node
	size int
}

func (q *queue) getSize() int {
	return q.size
}

func (q *queue) isEmpty() bool {
	return q.size == 0
}

func (q *queue) enQueue(node *node) {
	q.data = append(q.data, node)
	q.size++
}

func (q *queue) deQueue() *node {
	if q.isEmpty() {
		panic("deQueue failed. queue is empty")
	}
	ret := q.data[0]
	q.data = q.data[1:]
	q.size--
	return ret
}

func newQueue() *queue {
	return &queue{
		data: make([]*node, 0),
		size: 0,
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
		b.size++
		return newNode(e)
	}
	if e < n.val {
		n.left = b.add(n.left, e)
	} else if e > n.val {
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

// 非递归实现二分搜索树的前序遍历
func (b *BST) PreOrderNR() {
	stack := newStack()
	stack.push(b.root)
	for !stack.isEmpty() {
		cur := stack.pop()
		fmt.Printf("%d, ", cur.val)
		if cur.right != nil {
			stack.push(cur.right)
		}

		if cur.left != nil {
			stack.push(cur.left)
		}
	}
}

// 函数式编程
func (b *BST) traverseFunc(n *node, f func(n *node)) {
	if n == nil {
		return
	}
	b.traverseFunc(n.left, f)
	f(n)
	b.traverseFunc(n.right, f)
}

// 前序遍历递归
func (b *BST) traverse(n *node) {
	if n == nil {
		return
	}
	fmt.Printf("%d, ", n.val)
	b.traverse(n.left)
	b.traverse(n.right)
}

func (b *BST) MaxValueNR() int {
	if b.size == 0 {
		panic("BST is empty")
	}
	root := b.root
	for root.right != nil {
		root = root.right
	}
	return root.val
}

func (b *BST) Max() int {
	if b.size == 0 {
		panic("BST is empty")
	}
	return b.max(b.root).val
}

func (b *BST) max(node *node) *node {
	if node.right == nil {
		return node
	}
	return b.max(node.right)
}

func (b *BST) Minimum() int {
	if b.size == 0 {
		panic("BST is empty")
	}
	return b.minimum(b.root).val
}

func (b *BST) minimum(node *node) *node {
	if node.left == nil {
		return node
	}
	return b.minimum(node.left)
}

func (b *BST) RemoveMax() int {
	ret := b.Max()
	b.root = b.removeMax(b.root)
	return ret
}

func (b *BST) removeMax(node *node) *node {
	if node.right == nil {
		left := node.left
		node.left = nil
		b.size--
		return left
	}
	node.right = b.removeMax(node.right)
	return node
}

func (b *BST) RemoveMin() int {
	ret := b.Minimum()
	b.root = b.removeMin(b.root)
	return ret
}

func (b *BST) removeMin(node *node) *node {
	if node.left == nil {
		right := node.right
		node.right = nil
		b.size--
		return right
	}
	node.left = b.removeMin(node.left)
	return node
}

func (b *BST) MinValueNR() int {
	root := b.root
	for root.left != nil {
		root = root.left
	}
	return root.val
}

func (b *BST) Remove(e int) {
	b.root = b.remove(b.root, e)
}

// 删除以node为根的二分搜索树中值为e的节点, 递归算法
// 返回删除节点后的二分搜索树的根
func (b *BST) remove(node *node, e int) *node {
	if node == nil {
		return nil
	}

	if e < node.val {
		node.left = b.remove(node.left, e)
		return node
	} else if e > node.val {
		node.right = b.remove(node.right, e)
		return node
	} else {
		// 待删除的节点左子数为空的情况
		if node.left == nil {
			right := node.right
			node.right = nil
			b.size--
			return right
		}
		// 待删除元素右子数为空的情况
		if node.right == nil {
			left := node.left
			node.left = nil
			b.size--
			return left
		}

		// 待删除元素左右子树都不为空的情况
		// 找到比待删除节点大的最小节点，即待删除节点右子树的最小节点
		successor := b.minimum(node.right)
		// 用这个节点顶替待删除节点的位置
		successor.right = b.removeMin(node.right)
		successor.left = node.left
		node.left = nil
		node.right = nil
		return successor

	}
}

// 层序遍历
func (b *BST) LevelOrder() {
	queue := newQueue()
	queue.enQueue(b.root)
	for !queue.isEmpty() {
		cur := queue.deQueue()
		fmt.Printf("%d, ", cur.val)
		if cur.left != nil {
			queue.enQueue(cur.left)
		}
		if cur.right != nil {
			queue.enQueue(cur.right)
		}
	}
}

func (b *BST) Floor(e int) int {
	return b.floor(b.root, e).val
}

func (b *BST) floor(node *node, e int) *node {
	if node == nil {
		return nil
	}

	if e == node.val {
		return node
	}

	if e < node.val {
		return b.floor(node.left, e)
	}

	result := b.floor(node.right, e)
	if result != nil {
		return result
	}
	return node
}

func (b *BST) Rank(e int) int {
	count := 0
	result := 0
	b.traverseFunc(b.root, func(node *node) {
		count++
		if node.val == e {
			result = count
			return
		}
	})
	return result
}


func (b *BST) Select(rank int) (e int) {
	count := 0
	result := 0
	b.traverseFunc(b.root, func(node *node) {
		count++
		if count == rank {
			result = node.val
			return
		}
	})
	return result
}
