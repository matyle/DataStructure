package linear

import (
	"fmt"
)

const DEFAULT_CAP = 100

// Path: array.go
type Array[T any] struct {
	data []T
	len  int
}

func NewArray[T any](cap int) *Array[T] {
	return &Array[T]{
		len:  0,
		data: make([]T, cap), //全部为零值
	}
}

// impl the linear
func (a *Array[T]) Len() int {
	return a.len
}

func (a *Array[T]) Insert(index int, element T) bool {
	if index < 0 || index > a.len {
		fmt.Println("index out of range")
		return false
	}

	if a.len == len(a.data) {
		// 扩容
		a.data = append(a.data, make([]T, len(a.data))...)
	}

	for i := a.len; i > index; i-- {
		a.data[i] = a.data[i-1]
	}
	a.data[index] = element
	a.len++
	return true
}

func (a *Array[T]) Remove(index int) bool {
	if index < 0 || index > a.len {
		fmt.Println("index out of range")
		return false
	}
	//[0,1,2,3,4,5]
	for i := index; i < a.len-1; i++ {
		a.data[i] = a.data[i+1]
	}
	a.len--
	return true
}

func (a *Array[T]) Get(index int) T {
	var res T
	if index < 0 || index > a.len {
		fmt.Println("index out of range")
		return res
	}
	return a.data[index]
}

// Main
func ArrayMain() {
	arr := NewArray[int](DEFAULT_CAP)
	for i := 0; i < 100; i++ {
		arr.Insert(i, i)
	}
	for i := 0; i < 100; i++ {
		fmt.Printf("%d ", arr.Get(i))
	}
	fmt.Println()
	arr.Remove(0)
	for i := 0; i < arr.Len(); i++ {
		fmt.Printf("%d ", arr.Get(i))
	}
	fmt.Println()
	arr.Remove(9)
	for i := 0; i < arr.Len(); i++ {
		fmt.Printf("%d ", arr.Get(i))
	}
}
