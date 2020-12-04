package main

import "fmt"

// 请判断一个链表是否为回文链表。 https://leetcode-cn.com/problems/palindrome-linked-list

// 示例 1:
// 输入: 1->2
// 输出: false
// 示例 2:
// 输入: 1->2->2->1
// 输出: true
// 进阶：
// 你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

// 思路1：
// 快慢指针，慢指针反转，快指针控制进度
// 再次使用慢指针和从头开始遍历，并比较

func main() {
	l := NewList([]int{1, 2, 3, 4, 3, 2, 1})
	fmt.Println(isPalindrome(l))
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	slow, fast := head, head
	var pre *ListNode
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		tmpNext := slow.Next
		slow.Next = pre
		pre = slow
		slow = tmpNext
	}
	p1 := slow.Next
	p2 := pre
	for p1 != nil {
		// fmt.Println("p1:", p1.Val)
		// fmt.Println("p2:", p2.Val)
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}

func NewList(arr []int) *ListNode {
	head := &ListNode{}
	tail := head
	for _, i := range arr {
		tail.Next = &ListNode{Val: i}
		tail = tail.Next
	}
	return head.Next
}

func Print(head *ListNode) {
	if head == nil {
		fmt.Println("empty list")
	}
	for head != nil {
		fmt.Printf("%+v->", head.Val)
		head = head.Next
	}
	fmt.Printf("\n")
}
