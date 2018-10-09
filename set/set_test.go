package set

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	bst := newBst()
	bst.add("b")
	bst.add("c")
	bst.add("d")
	bst.add("a")
	fmt.Println(bst)
	fmt.Println("max", bst.maximum())
	fmt.Println("min", bst.minimum())

	fmt.Printf("移除最大元素%s后二叉树:%s\n", bst.removeMaximum(), bst)
	fmt.Printf("移除最小元素%s后二叉树:%s\n", bst.removeMinimum(), bst)

}

func TestBSTSet(t *testing.T) {
	bytes, err := ioutil.ReadFile("/Users/xiaozefeng/Desktop/a-tale-of-two-cities.txt")
	if err != nil {
		panic(err)
	}
	fields := strings.Fields(string(bytes))
	set := New()
	for _, v := range fields {
		set.Add(v)
	}
	fmt.Println("元素格式:", set.Size())
}
