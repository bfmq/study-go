package _8_stack

import "fmt"

type Browser struct {
	forwardStack Stack
	backStack    Stack
}

func NewBrowser() *Browser {
	return &Browser{
		forwardStack: NewArrayStack(),
		backStack:    NewLinkedListStack(),
	}
}

func (b *Browser) CanForward() bool {
	if b.forwardStack.IsEmpty() {
		return false
	}
	return true
}

func (b *Browser) CanBack() bool {
	if b.backStack.IsEmpty() {
		return false
	}
	return true
}

// 网页前进
func (b *Browser) Forward() {
	if b.forwardStack.IsEmpty() {
		return
	}
	v := b.forwardStack.Pop()
	b.backStack.Push(v)
	fmt.Printf("forward to %+v\n", v)
}

// 网页后退
func (b *Browser) Back() {
	if b.backStack.IsEmpty() {
		return
	}
	v := b.backStack.Pop()
	b.forwardStack.Push(v)
	fmt.Printf("back to %+v\n", v)
}
