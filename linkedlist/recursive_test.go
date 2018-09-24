package linkedlist

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8}
	res := sum(sl, 0)
	fmt.Println(res)
}

func sum(sl []int, l int) int {
	if l == len(sl) {
		return 0
	}
	return sl[l] + sum(sl, l+1)
}
