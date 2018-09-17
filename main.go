package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr1 := make([]int, 0)
	for i := 0; i < 10; i++ {
		arr1 = append(arr1, i)
	}
	arr2 := make([]int, 0)
	arr3 := make([]int, 0)
	fmt.Printf("%+v\n", arr1)
	fmt.Printf("%#v\n", arr2)
	fmt.Printf("%T\n", arr3)
	fmt.Println(strconv.IntSize)
}
