package main

import (
	"fmt"
	"math"
)

type MinStack struct {
	// 数据栈
	dataList []int
	// 最小栈
	minList []int
}

func Constructor() MinStack {
	return MinStack{
		dataList: []int{},
		// 使用math.MaxInt64方便第一个值的入栈
		minList: []int{math.MaxInt64},
	}
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func (this *MinStack) Push(val int) {
	// 数据栈正常添加
	this.dataList = append(this.dataList, val)
	// 每次添加值，最小栈添加（当前栈内最小值与当前值）比较后的更小值
	this.minList = append(this.minList, min(val, this.minList[len(this.minList)-1]))
}

func (this *MinStack) Pop() {
	// 数据栈正常减少
	this.dataList = this.dataList[:len(this.dataList)-1]
	// 每次减少值，最小栈也减一，因为每个值在添加时对应的最小值已经加入了最小栈，是一一对应关系
	this.minList = this.minList[:len(this.minList)-1]
}

func (this *MinStack) Top() int {
	// 数据栈最后的值
	return this.dataList[len(this.dataList)-1]
}

func (this *MinStack) GetMin() int {
	// 最小栈最后的值
	return this.minList[len(this.minList)-1]
}

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	obj.Pop()
	param_3 := obj.Top()
	param_4 := obj.GetMin()
	fmt.Println(param_3)
	fmt.Println(param_4)
}
