package linkedlist

type RecursiveLinkedList struct {
	dummyHead *node
	size      int
}

func NewRecursiveLinkedList() *RecursiveLinkedList {
	return &RecursiveLinkedList{
		dummyHead: newNode(-1),
		size:      0,
	}
}

func (ls *RecursiveLinkedList) Size() int {
	return ls.size
}
func (ls *RecursiveLinkedList) IsEmpty() bool {
	return ls.size == 0
}

func (ls *RecursiveLinkedList) Add(index, e int) {
	// todo
}
