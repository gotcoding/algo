package _2_linkedlist

import "fmt"

/*
单链表操作
*/
type Node struct {
	next  *Node
	value interface{}
}

type LinkedList struct {
	head   *Node
	length uint
}

func NewNode(v interface{}) *Node {
	return &Node{nil, v}
}

func (this *Node) GetNext() *Node {
	return this.next
}

func (this *Node) GetValue() interface{} {
	return this.value
}

// 创建链表
func NewLinkedList() *LinkedList {
	return &LinkedList{NewNode(0), 0}
}

// 在某个节点前面插入节点
func (this *LinkedList) InsertBefore(p *Node, v interface{}) bool {
	if p == nil || p == this.head {
		return false
	}
	cur := this.head.next
	pre := this.head
	for cur != nil {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if cur == nil {
		return false
	}
	newNode := NewNode(v)
	pre.next = newNode
	newNode.next = cur
	this.length++
	return true
}

// 在某个节点后面插入节点
func (this *LinkedList) InsertAfter(p *Node, v interface{}) bool {
	if p == nil {
		return false
	}
	newNode := NewNode(v)
	oldNext := p.next
	newNode.next = oldNext
	p.next = newNode
	this.length++
	return true
}

// 在链表头部插入节点
func (this *LinkedList) InsertToHead(v interface{}) bool {
	return this.InsertAfter(this.head, v)
}

// 在链表尾部插入节点
func (this *LinkedList) InsertToTail(v interface{}) bool {
	cur := this.head
	for cur != nil {
		cur = cur.next
	}
	return this.InsertAfter(cur, v)
}

// 通过索引查找节点
func (this *LinkedList) FindByIndex(index uint) *Node {
	if index > this.length {
		return nil
	}
	cur := this.head.next
	var i uint = 0
	for ; i < index; i++ {
		cur = cur.next
	}
	return cur
}

// 删除传入的节点
func (this *LinkedList) Delete(p *Node) bool {
	if p == nil {
		return false
	}
	cur := this.head.next
	pre := this.head
	for nil != cur {
		if p == cur {
			break
		}
		pre = cur
		cur = cur.next
	}
	if cur == nil {
		return false
	}
	pre.next = cur.next
	this.length--
	return true
}

// 打印链表
func (this *LinkedList) Print() {
	cur := this.head.next
	format := ""
	for cur != nil {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if cur != nil {
			format += "->"
		}
	}
	fmt.Println(format)
}
