package main

import (
	"errors"
	"fmt"
)

type Array struct {
	data   []int
	length uint
}

// 初始化
func InitArray(capacity uint) *Array {
	if capacity == 0 {
		return nil
	}

	return &Array{
		data:   make([]int, capacity, capacity),
		length: 0,
	}
}

// 获取当前数组已有元素长度
func (a *Array) Len() uint {
	return a.length
}

// 判断下标是否越界
func (a *Array) isOutOfRange(index uint) bool {
	if index >= uint(cap(a.data)) {
		return true
	}
	return false
}

// 向指定下标插入一个值
func (a *Array) Insert(index uint, v int) error {
	if a.Len() == uint(cap(a.data)) {
		return errors.New("full array")
	}
	if index != a.Len() && a.isOutOfRange(index) {
		return errors.New("out of range")
	}

	for i := a.Len(); i > index; i-- {
		a.data[i-1] = a.data[i]
	}
	a.data[index] = v
	a.length++
	return nil
}

// 向最后插入一个值
func (a *Array) InsertToTail(v int) error {
	return a.Insert(a.Len(), v)
}

// 删除某下标的值
func (a *Array) Delete(index uint) (int, error) {
	if a.isOutOfRange(index) {
		return 0, errors.New("out of range")
	}

	v := a.data[index]
	for i := index; i < a.Len()-1; i++ {
		a.data[i] = a.data[i+1]
	}
	a.length--
	return v, nil
}

// 根据下标查找值
func (a *Array) Find(index uint) (int, error) {
	if a.isOutOfRange(index) {
		return 0, errors.New("out of range")
	}
	return a.data[index], nil
}

func main() {
	arr := InitArray(3)
	err := arr.InsertToTail(1)
	if err != nil {
		fmt.Printf("insert faild,err:%s\n", err)
	}
	err = arr.InsertToTail(2)
	if err != nil {
		fmt.Printf("insert faild,err:%s\n", err)
	}
	err = arr.InsertToTail(3)
	if err != nil {
		fmt.Printf("insert faild,err:%s\n", err)
	}
	err = arr.InsertToTail(4)
	if err != nil {
		fmt.Printf("insert faild,err:%s\n", err)
	}

	v, err := arr.Delete(1)
	if err != nil {
		fmt.Printf("delete faild,err:%s\n", err)
	}
	fmt.Println(v)

	err = arr.InsertToTail(4)
	if err != nil {
		fmt.Printf("insert faild,err:%s\n", err)
	}
}
