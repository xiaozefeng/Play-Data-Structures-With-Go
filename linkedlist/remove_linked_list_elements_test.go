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
	n := removeElements2(ln, 2)
	fmt.Println(n)

}
