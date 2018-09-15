package array

import (
	"errors"
	"fmt"
)

type Array struct {
	data []int
	size int
}

func (arr *Array) String() string {
	result := ""
	result = fmt.Sprintf("Array: len:%d, size:%d ", len(arr.data), arr.size)
	result += fmt.Sprintf(", Data: ")
	for i := 0; i < arr.size; i++ {
		result += fmt.Sprintf("%d ", arr.data[i])
	}
	return result
}

// the factory method
func NewArrayWithCapacity(capacity int) *Array {
	return &Array{
		data: make([]int, capacity, capacity),
		size: 0,
	}
}

func NewArray() *Array {
	return NewArrayWithCapacity(10)
}

func (arr *Array) Size() int {
	return arr.size
}

func (arr *Array) IsEmpty() bool {
	return arr.size == 0
}

// 添加一个元素
func (arr *Array) Add(index, e int) (err error) {
	if arr.size == len(arr.data) {
		arr.resize(len(arr.data) * 2)
	}
	if index < 0 || index > arr.size {
		return errors.New("index required >= 0 and <= size")
	}
	for i := arr.size; i > index; i-- {
		arr.data[i] = arr.data[i-1]
	}
	arr.data[index] = e
	arr.size++
	return nil
}

// 在尾部添加一个元素
func (arr *Array) AddLast(e int) error {
	return arr.Add(arr.size, e)
}

//在头部添加一个元素
func (arr *Array) AddFirst(e int) error {
	return arr.Add(0, e)
}

// 根据索引删除一个元素
func (arr *Array) Remove(index int) (e int, err error) {
	if index < 0 || index >= arr.size {
		return 0, errors.New("index required >= 0 and < size")
	}
	ret := arr.data[index]
	for i := index; i < arr.size; i++ {
		arr.data[i] = arr.data[i+1]
	}
	arr.size--
	// 只有四分之一的空间在用的时候，减少容量
	if len(arr.data)/4 == arr.size && len(arr.data)/2 != 0 {
		arr.resize(len(arr.data) / 2)
	}
	return ret, nil

}

func (arr *Array) RemoveFirst() (e int, err error) {
	return arr.Remove(0)
}

func (arr *Array) RemoveLast() (e int, err error) {
	return arr.Remove(arr.size - 1)
}

func (arr *Array) Contains(e int) bool {
	for _, val := range arr.data {
		if val == e {
			return true
		}
	}
	return false
}

// 查找并删除第一个元素
func (arr *Array) RemoveElement(e int) bool {
	index := arr.Find(e)
	if index != -1 {
		arr.Remove(index)
		return true
	}
	return false
}

// 修改一个元素
func (arr *Array) Set(index, e int) (err error) {
	if index < 0 || index >= arr.size {
		return errors.New("index required >= 0 and <size")
	}
	arr.data[index] = e
	return nil
}

// 根据元素查询索引
func (arr *Array) Find(e int) (index int) {
	for index, val := range arr.data {
		if val == e {
			return index
		}
	}
	return -1
}

func (arr *Array) resize(capacity int) {
	newData := make([]int, capacity, capacity)
	for i := 0; i < arr.size; i++ {
		newData[i] = arr.data[i]
	}
	arr.data = newData
}
func (arr *Array) Get(index int) (e int, err error) {
	if index < 0 || index >= arr.size {
		return -1, errors.New("index required >0 and < size")
	}
	return arr.data[index], nil
}

func (arr *Array) GetLast() (e int, err error) {
	return arr.Get(arr.size - 1)
}
