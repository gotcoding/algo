package stack

import "fmt"

/*
模拟浏览器前进和倒退功能
*/

// Browser 浏览器
type Browser struct {
	forwardStack Stack
	backStack    Stack
}

// NewBrowser 新建浏览器
func NewBrowser() *Browser {
	return &Browser{
		forwardStack: NewArrayStack(),
		backStack:    NewLinkedListStack(),
	}
}

// CanForward 是否可以前进
func (b *Browser) CanForward() bool {
	if b.forwardStack.IsEmpty() {
		return false
	}
	return true
}

// CanBack 是否可以前进
func (b *Browser) CanBack() bool {
	if b.backStack.IsEmpty() {
		return false
	}
	return true
}

// Open 打开浏览器
func (b *Browser) Open(addr string) {
	fmt.Printf("Open new addr %+v\n", addr)
	b.forwardStack.Flush()
}

// NewPage 浏览页面
func (b *Browser) NewPage(addr string) {
	b.backStack.Push(addr)
}

// Forward 向前
func (b *Browser) Forward() {
	if b.forwardStack.IsEmpty() {
		return
	}
	top := b.forwardStack.Pop()
	b.backStack.Push(top)
	fmt.Printf("forward to %+v\n", top)
}

// Back 向后
func (b *Browser) Back() {
	if b.backStack.IsEmpty() {
		return
	}
	top := b.backStack.Pop()
	b.forwardStack.Push(top)
	fmt.Printf("back to %+v\n", top)
}
