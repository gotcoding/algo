package main

// 92. 反转链表 II https://leetcode-cn.com/problems/reverse-linked-list-ii
// 反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

// 说明:
// 1 ≤ m ≤ n ≤ 链表长度。

// 示例:
// 输入: 1->2->3->4->5->NULL, m = 2, n = 4
// 输出: 1->4->3->2->5->NULL

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	prem := dummy
	for i := 1; i <= m-1; i++ {
		prem = prem.Next
	}
	// 在m和n之间反转
	cur := prem.Next
	var pre *ListNode
	for i := m; i <= n; i++ {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	prem.Next.Next = cur
	prem.Next = pre
	return dummy.Next
}

func main() {

}
