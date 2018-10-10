package set

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	bst := newBst()
	bst.add("b")
	bst.add("c")
	bst.add("d")
	bst.add("a")
	bst.add("a")
	fmt.Println(bst)
	fmt.Println("max", bst.maximum())
	fmt.Println("min", bst.minimum())

	fmt.Printf("移除最大元素%s后二叉树:%s\n", bst.removeMaximum(), bst)
	fmt.Printf("移除最小元素%s后二叉树:%s\n", bst.removeMinimum(), bst)

}

func TestBSTSet(t *testing.T) {
	set := New()
	file, err := os.Open("/Users/xiaozefeng/Desktop/a-tale-of-two-cities.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fields := strings.Fields(string(line))
		for _, v := range fields {
			set.Add(v)
		}
	}
	fmt.Println("bst size:", set.Size())

}
