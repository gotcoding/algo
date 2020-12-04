package dev

import "fmt"

// 链表常见操作
// 新建链表
// 打印链表
// 判断链表是否有环
// 合并两个有序的链表
// 反转链表
// 部分反转链表
// 按照组反转链表

type ListNode struct {
	Value int
	Next  *ListNode
}

func NewList(arr []int) *ListNode {
	head := &ListNode{}
	tail := head
	for _, i := range arr {
		tail.Next = &ListNode{Value: i}
		tail = tail.Next
	}
	return head.Next
}

func Print(head *ListNode) {
	if head == nil {
		fmt.Println("empty list")
	}
	for head != nil {
		fmt.Printf("%+v->", head.Value)
		head = head.Next
	}
	fmt.Printf("\n")
}

func hasCircle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// 合并两个递增的链表
func mergeList(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tail := head
	for l1 != nil && l2 != nil {
		if l1.Value > l2.Value {
			tail.Next = &ListNode{Value: l2.Value}
			l2 = l2.Next
		} else {
			tail.Next = &ListNode{Value: l1.Value}
			l1 = l1.Next
		}
		tail = tail.Next
	}
	if l1 == nil {
		tail.Next = l2
	}
	if l2 == nil {
		tail.Next = l1
	}
	return head.Next
}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		tmpNext := head.Next
		head.Next = pre
		pre = head
		head = tmpNext
	}
	return pre
}

func reverseBetween(head *ListNode, m, n int) *ListNode {
	// 增加哨兵节点，简化head处理
	guard := &ListNode{Next: head}
	prem := guard
	// 跳转到m-1
	for i := 1; i <= m-1; i++ {
		prem = prem.Next
	}
	// 反转m到n
	var pre *ListNode
	cur := prem.Next
	for i := m; i <= n; i++ {
		tmpNext := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmpNext
	}
	// 缝合链表
	prem.Next.Next = cur
	prem.Next = pre
	return guard.Next
}

func reverseGroupK(head *ListNode, k int) *ListNode {
	// 哨兵节点
	guard := &ListNode{Next: head}
	pre := guard
	for head != nil {
		tail := head
		for i := 1; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return guard.Next
			}
		}
		head, tail = reverseNodes(head, tail)
		pre.Next = head
		head = tail.Next
		pre = tail
	}
	return guard.Next
}

func reverseNodes(head, tail *ListNode) (*ListNode, *ListNode) {
	cur := head
	pre := tail.Next
	for pre != tail {
		tmpNext := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmpNext
	}
	return tail, head
}
