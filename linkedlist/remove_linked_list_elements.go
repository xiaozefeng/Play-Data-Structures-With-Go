package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	//删除头节点特殊处理
	for head != nil && head.Val == val {
		head = head.Next
	}
	// 整个链表都删除空了
	if head == nil {
		return nil
	}
	pre := head
	for pre.Next != nil {
		if pre.Next.Val == val {
			pre.Next = pre.Next.Next
		} else {
			pre = pre.Next
		}
	}
	return head
}

func removeElements2(head *ListNode, val int) *ListNode {
	// 创建虚拟头结点
	dummyHead := &ListNode{
		Val:  -1,
		Next: head,
	}
	pre := dummyHead
	for pre.Next != nil {
		if pre.Next.Val == val {
			pre.Next = pre.Next.Next
		} else {
			pre = pre.Next
		}
	}
	return dummyHead.Next

}

func removeElements3(head *ListNode, val int) *ListNode {
	// 最基本的问题
	if head == nil {
		return nil
	}
	head.Next = removeElements3(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}
