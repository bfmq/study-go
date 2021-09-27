package _8_stack

import "fmt"

type ArrayStack struct {
	data []interface{}
	top  int
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, 32),
		top:  -1,
	}
}

// 验证是否为空
func (as *ArrayStack) isEmpty() bool {
	if as.top < 0 {
		return true
	}
	return false
}

// 压栈
func (as *ArrayStack) Push(v interface{}) {
	if as.top < 0 {
		as.top = 0
	} else {
		as.top++
	}

	as.data = append(as.data, v)
}

// 出栈
func (as *ArrayStack) Pop() interface{} {
	if as.isEmpty() {
		return nil
	}
	v := as.data[as.top]
	as.top--
	return v

}

// 获取栈顶元素
func (as *ArrayStack) Top(v interface{}) interface{} {
	if as.isEmpty() {
		return nil
	}
	return as.data[as.top]
}

// 清空栈
func (as *ArrayStack) Flush() {
	as.top = -1
}

// 打印
func (as *ArrayStack) Print() {
	if as.isEmpty() {
		fmt.Println("empty stack")
	} else {
		for i := as.top; i >= 0; i-- {
			println(as.data[i])
		}
	}
}
