package main

type MyCircularDeque struct {
	data   []int
	head   int
	tail   int
	size   int
	length int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		data:   make([]int, k),
		head:   0,
		tail:   0,
		size:   k,
		length: 0,
	}
}

// /** Adds an item at the front of Deque. Return true if the operation is successful. */
// func (this *MyCircularDeque) InsertFront(value int) bool {
// 	if this.IsFull() {
// 		return false
// 	}
// 	if this.length == 0 {
// 		this.data[this.head] = value
// 	} else {
// 		pre := (this.head - 1 + this.size) % this.size
// 		this.data
// 	}
// 	this.length++
// 	return true
// }

// /** Adds an item at the rear of Deque. Return true if the operation is successful. */
// func (this *MyCircularDeque) InsertLast(value int) bool {

// }

// /** Deletes an item from the front of Deque. Return true if the operation is successful. */
// func (this *MyCircularDeque) DeleteFront() bool {

// }

// /** Deletes an item from the rear of Deque. Return true if the operation is successful. */
// func (this *MyCircularDeque) DeleteLast() bool {

// }

// /** Get the front item from the deque. */
// func (this *MyCircularDeque) GetFront() int {

// }

// /** Get the last item from the deque. */
// func (this *MyCircularDeque) GetRear() int {

// }

// /** Checks whether the circular deque is empty or not. */
// func (this *MyCircularDeque) IsEmpty() bool {
// 	return this.length == 0
// }

// /** Checks whether the circular deque is full or not. */
// func (this *MyCircularDeque) IsFull() bool {
// 	return this.length == this.size
// }
