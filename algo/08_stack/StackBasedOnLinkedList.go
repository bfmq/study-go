package _8_stack

import "fmt"

type Node struct {
	next *Node
	val  interface{}
}

type LinkedListStack struct {
	topNode *Node
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{
		nil,
	}
}

// 验证是否为空
func (lls *LinkedListStack) isEmpty() bool {
	return lls.topNode == nil
}

// 压栈
func (lls *LinkedListStack) Push(v interface{}) {
	lls.topNode = &Node{lls.topNode, v}
}

// 出栈
func (lls *LinkedListStack) Pop() interface{} {
	if lls.isEmpty() {
		return nil
	}
	v := lls.topNode.val
	lls.topNode = lls.topNode.next
	return v

}

// 获取栈顶元素
func (lls *LinkedListStack) Top(v interface{}) interface{} {
	if lls.isEmpty() {
		return nil
	}
	return lls.topNode.val
}

// 清空栈
func (lls *LinkedListStack) Flush() {
	lls.topNode = nil
}

// 打印
func (lls *LinkedListStack) Print() {
	if lls.isEmpty() {
		fmt.Println("empty stack")
	} else {
		for i := lls.topNode; i != nil; i = i.next {
			println(i.val)
		}
	}
}
