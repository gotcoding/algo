package queue

import "fmt"

// Node 节点
type Node struct {
	val  interface{}
	next *Node
}

// LinkedListQueue 链表队列
type LinkedListQueue struct {
	head   *Node
	tail   *Node
	length int
}

// NewNode 新建节点
func NewNode(v interface{}) *Node {
	return &Node{v, nil}
}

// NewLinkedListQueue 新建队列
func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{nil, nil, 0}
}

// EnQueue 入队
func (q *LinkedListQueue) EnQueue(v interface{}) {
	n := NewNode(v)
	if q.head == nil {
		q.head = n
		q.tail = n
	}
	q.tail.next = n
	q.tail = n
	q.length++
}

// DeQueue 出队
func (q *LinkedListQueue) DeQueue() interface{} {
	if q.head == nil {
		return nil
	}
	v := q.head.val
	q.head = q.head.next
	q.length--
	return v
}

// Print 打印
func (q *LinkedListQueue) Print() {
	if q.head == nil {
		fmt.Println("empty queue")
	}
	result := "head"
	for cur := q.head; cur != nil; cur = cur.next {
		result += fmt.Sprintf("<-%+v", cur.val)
	}
	result += "<-tail"
	fmt.Println(result)
}
