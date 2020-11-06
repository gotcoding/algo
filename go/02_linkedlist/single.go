package linkedlist

import "fmt"

/*
单链表操作
链表都有哪些操作？
1.创建一个链表；
2.插入数据：在某个节点前，在某个节点后，在链头插入，在链尾插入；
3.删除数据；
4.查找数据；
5.打印链表；
6.单链表反转
7.判断链表是否有环；
*/
type LinkedNode struct {
	next  *LinkedNode
	value interface{}
}

type LinkedList struct {
	head   *LinkedNode
	length uint
}

func NewNode(v interface{}) *LinkedNode {
	return &LinkedNode{nil, v}
}

func (ln *LinkedNode) GetNext() *LinkedNode {
	return ln.next
}

func (ln *LinkedNode) GetValue() interface{} {
	return ln.value
}

//NewLinkedList 创建链表
func NewLinkedList() *LinkedList {
	return &LinkedList{NewNode(0), 0}
}

// InsertBefore 在某个节点前面插入节点
func (ll *LinkedList) InsertBefore(p *LinkedNode, v interface{}) bool {
	if p == nil || p == ll.head {
		return false
	}
	cur := ll.head.next
	pre := ll.head
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
	ll.length++
	return true
}

// InsertAfter 在某个节点后面插入节点
func (ll *LinkedList) InsertAfter(p *LinkedNode, v interface{}) bool {
	if p == nil {
		return false
	}
	newNode := NewNode(v)
	oldNext := p.next
	newNode.next = oldNext
	p.next = newNode
	ll.length++
	return true
}

// InsertToHead 在链表头部插入节点
func (ll *LinkedList) InsertToHead(v interface{}) bool {
	return ll.InsertAfter(ll.head, v)
}

// InsertToTail 在链表尾部插入节点
func (ll *LinkedList) InsertToTail(v interface{}) bool {
	cur := ll.head
	for cur.next != nil {
		cur = cur.next
	}
	return ll.InsertAfter(cur, v)
}

// FindByIndex 通过索引查找节点
func (ll *LinkedList) FindByIndex(index uint) *LinkedNode {
	if index > ll.length {
		return nil
	}
	cur := ll.head.next
	var i uint = 0
	for ; i < index; i++ {
		cur = cur.next
	}
	return cur
}

// Delete 删除传入的节点
func (ll *LinkedList) Delete(p *LinkedNode) bool {
	if p == nil {
		return false
	}
	cur := ll.head.next
	pre := ll.head
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
	ll.length--
	return true
}

// Print 打印链表
func (ll *LinkedList) Print() {
	cur := ll.head.next
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

// Reverse 链表反转 时间复杂度：O(N)
func (ll *LinkedList) Reverse() {
	if ll.head == nil || ll.head.next == nil || ll.head.next.next == nil {
		return
	}
	var pre *LinkedNode = nil
	cur := ll.head.next
	for cur != nil {
		tmp := cur.next
		cur.next = pre
		pre = cur
		cur = tmp
	}
	ll.head.next = pre
}

// HasCycle 判断是否有环
// 思路：快慢指针，慢指针每次走一步，快指针每次走两步，如果存在环的话，两个指针必定会相遇
func (ll *LinkedList) HasCycle() bool {
	if ll.head != nil {
		slow := ll.head
		fast := ll.head
		for fast != nil && fast.next != nil {
			slow = slow.next
			fast = fast.next.next
			if slow == fast {
				return true
			}
		}
	}
	return false
}

// FindMiddleNode 获取中间节点
// 思路：使用快慢指针
func (ll *LinkedList) FindMiddleNode() *LinkedNode {
	if ll.head == nil || ll.head.next == nil {
		return nil
	}
	if nil == ll.head.next.next {
		return ll.head.next
	}
	slow, fast := ll.head, ll.head
	for nil != fast && nil != fast.next {
		slow, fast = slow.next, fast.next.next
	}
	return slow
}

// MergeSortedList 合并两个有序的链表
func MergeSortedList(l1, l2 *LinkedList) *LinkedList {
	if nil == l1 || nil == l1.head || nil == l1.head.next {
		return l2
	}
	if nil == l2 || nil == l2.head || nil == l2.head.next {
		return l1
	}
	l := NewLinkedList()
	cur := l.head
	cur1 := l1.head.next
	cur2 := l2.head.next
	for nil != cur1 && nil != cur2 {
		if cur1.value.(int) > cur2.value.(int) {
			cur.next = cur2
			cur2 = cur2.next
		} else {
			cur.next = cur1
			cur1 = cur1.next
		}
		cur = cur.next
	}
	if nil != cur1 {
		cur.next = cur1
	} else if nil != cur2 {
		cur.next = cur2
	}

	return l
}

// DeleteBottomN 删除倒数第n个节点
func (ll *LinkedList) DeleteBottomN(n int) {
	if n <= 0 || nil == ll.head || nil == ll.head.next {
		return
	}

	fast := ll.head
	for i := 1; i <= n && fast != nil; i++ {
		fast = fast.next
	}

	if nil == fast {
		return
	}

	slow := ll.head
	for nil != fast.next {
		slow = slow.next
		fast = fast.next
	}
	slow.next = slow.next.next
}
