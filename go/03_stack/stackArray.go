package stack

import "fmt"

/* 基于数组实现的栈 */

//ArrayStack 数组实现的栈
type ArrayStack struct {
	data []interface{} // 数据
	top  int           // 栈顶指针
}

// NewArrayStack 创建栈
func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, 32),
		top:  -1,
	}
}

// Push 压栈
func (as *ArrayStack) Push(v interface{}) {
	as.data = append(as.data, v)
	as.top++
}

// Print 打印
func (as *ArrayStack) Print() {
	if as.top == -1 {
		fmt.Println("empty stack")
	}
	for i := as.top; i >= 0; i-- {
		fmt.Println(as.data[i])
	}
}

// Pop 出栈
func (as *ArrayStack) Pop() interface{} {
	if as.top == -1 {
		return nil
	}
	v := as.data[as.top]
	as.top--
	return v
}

// IsEmpty 是否为空
func (as *ArrayStack) IsEmpty() bool {
	if as.top < 0 {
		return true
	}
	return false
}

// Top 获取栈顶数据
func (as *ArrayStack) Top() interface{} {
	if as.IsEmpty() {
		return nil
	}
	return as.data[as.top]
}

// Flush 清空栈
func (as *ArrayStack) Flush() {
	as.top = -1
}
