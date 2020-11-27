package main

import "fmt"

// 设计实现双端队列。
// 你的实现需要支持以下操作：
// MyCircularDeque(k)：构造函数,双端队列的大小为k。
// insertFront()：将一个元素添加到双端队列头部。 如果操作成功返回 true。
// insertLast()：将一个元素添加到双端队列尾部。如果操作成功返回 true。
// deleteFront()：从双端队列头部删除一个元素。 如果操作成功返回 true。
// deleteLast()：从双端队列尾部删除一个元素。如果操作成功返回 true。
// getFront()：从双端队列头部获得一个元素。如果双端队列为空，返回 -1。
// getRear()：获得双端队列的最后一个元素。 如果双端队列为空，返回 -1。
// isEmpty()：检查双端队列是否为空。
// isFull()：检查双端队列是否满了。

// 链表实现
type MyCircularDeque struct {
	head *Node
	tail *Node
	len  int
	size int
}

type Node struct {
	last  *Node
	next  *Node
	value int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		head: nil,
		tail: nil,
		size: k,
		len:  0,
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	node := &Node{nil, nil, value}
	if this.head == nil {
		this.head = node
		this.tail = node
	} else {
		this.head.last = node
		node.next = this.head
		this.head = node
	}
	this.len++
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	node := &Node{nil, nil, value}
	if this.head == nil {
		this.head = node
		this.tail = node
	} else {
		this.tail.next = node
		node.last = this.tail
		this.tail = node
	}
	this.len++
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	if this.head == this.tail {
		this.head = nil
		this.tail = nil
	} else {
		this.head = this.head.next
	}
	this.len--
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	if this.head == this.tail {
		this.head = nil
		this.tail = nil
	} else {
		this.tail = this.tail.last
	}
	this.len--
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.head.value
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.tail.value
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.head == nil
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.size == this.len
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */

func main() {
	circularDeque := Constructor(3) // 设置容量大小为3
	circularDeque.InsertLast(1)     // 返回 true
	circularDeque.InsertLast(2)     // 返回 true
	circularDeque.InsertFront(3)    // 返回 true
	// 已经满了，返回 false
	fmt.Println(circularDeque.InsertFront(4))
	// 返回 2
	fmt.Println(circularDeque.GetRear())
	fmt.Println(circularDeque)
	circularDeque.IsFull()       // 返回 true
	circularDeque.DeleteLast()   // 返回 true
	circularDeque.InsertFront(4) // 返回 true
	circularDeque.GetFront()     // 返回 4
}
