package main

// K 个一组翻转链表 https://leetcode-cn.com/problems/reverse-nodes-in-k-group
// 给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
// k 是一个正整数，它的值小于或等于链表的长度。
// 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

// 示例：
// 给你这个链表：1->2->3->4->5
// 当 k = 2 时，应当返回: 2->1->4->3->5
// 当 k = 3 时，应当返回: 3->2->1->4->5

func main() {

}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	// 哨兵节点
	guard := &ListNode{Next: head}
	pre := guard
	// 遍历链表
	for head != nil {
		// 找到尾节点
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil { //如果长度不够，直接跳出
				return guard.Next
			}
		}
		head, tail = reverse(head, tail)
		pre.Next = head
		// 为下次反转准备
		pre = tail
		head = tail.Next
	}
	return guard.Next
}

// 反转一个链表
func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		tmpNext := p.Next
		p.Next = prev
		prev = p
		p = tmpNext
	}
	return tail, head
}
