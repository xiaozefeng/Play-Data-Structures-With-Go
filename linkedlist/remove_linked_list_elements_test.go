package linkedlist

import (
	"fmt"
	"testing"
)

func Test_removeElements2(t *testing.T) {
	ln := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}
	n := removeElements3(ln, 1)
	fmt.Println(n)

}
