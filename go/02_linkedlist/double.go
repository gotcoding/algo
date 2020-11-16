package linkedlist

import "fmt"

type DoubleLinkedNode struct {
	Value interface{}
	Next  *DoubleLinkedNode
	Pre   *DoubleLinkedNode
}

type DoubleLinked struct {
	head   *DoubleLinkedNode
	length int
}

func NewDoubleLinked() *DoubleLinked {

	return &DoubleLinked{
		head:   &DoubleLinkedNode{},
		length: 0,
	}
}

func (dl *DoubleLinked) InsertToTail(v interface{}) {
	node := &DoubleLinkedNode{v, nil, nil}
	cur := dl.head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = node
	node.Pre = cur
	dl.length++
}

//Reverse 链表反转
func (dl *DoubleLinked) Reverse() {
	if dl.head == nil || dl.head.Next == nil || dl.head.Next.Next == nil {
		return
	}
	var pre, next *DoubleLinkedNode
	cur := dl.head.Next
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		cur.Pre = next
		pre = cur
		cur = next
	}
	dl.head.Next = pre
	pre.Pre = dl.head
}

func (dl *DoubleLinked) Print() {
	if dl.head == nil {
		fmt.Println("empty linked")
	}
	cur := dl.head.Next
	for cur != nil {
		fmt.Printf("->%v", cur.Value)
		cur = cur.Next
	}
	fmt.Printf("\n")
}
