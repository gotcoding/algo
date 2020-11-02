package stack

import "fmt"

/* 基于链表实现的栈 */

type node struct {
	next *node
	data interface{}
}

// LinkedListStack 链表实现的栈
type LinkedListStack struct {
	head *node
}

// newNode 新建节点
func newNode(v interface{}) *node {
	return &node{nil, v}
}

// NewLinkedListStack 新建栈
func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{newNode(0)}
}

// Push 压栈
func (s *LinkedListStack) Push(v interface{}) {
	n := newNode(v)
	if s.head.next == nil {
		s.head.next = n
	} else {
		n.next = s.head.next
		s.head.next = n
	}
}

// Print 打印栈
func (s *LinkedListStack) Print() {
	if s.head.next == nil {
		fmt.Println("empty stack")
	}
	cur := s.head.next
	for cur != nil {
		fmt.Println(cur.data)
		cur = cur.next
	}
}

// Pop 出栈
func (s *LinkedListStack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	v := s.head.next.data
	s.head.next = s.head.next.next
	return v
}

// IsEmpty 是否为空栈
func (s *LinkedListStack) IsEmpty() bool {
	if s.head.next == nil {
		return true
	}
	return false
}

// Top 获取栈顶数据
func (s *LinkedListStack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.head.next.data
}

// Flush 清空栈
func (s *LinkedListStack) Flush() {
	s.head.next = nil
}
