package reverse

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 迭代方法实现反转
func reverseIteration(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode // 用于记录前一个节点
	for head != nil { // 遍历链表
		tmpNext := head.Next // 记录后一个节点
		head.Next = pre      // 掉头反转
		pre = head           // pre指针前移一位
		head = tmpNext       // 遍历head指针后移一位
	}
	return pre
}

// 对链表m到n上进行反转
// 假设初始为
// 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7
// m=2，n=5
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	// 构建dummy，解决头结点问题
	dummy := &ListNode{Next: head}
	// 跳转到m-1节点，记录为prem
	prem := dummy
	for i := 1; i <= m-1; i++ {
		prem = prem.Next
	}
	cur := prem.Next // 跳转到m上，记录为cur
	// 对m到n进行反转
	var pre *ListNode
	for i := m; i <= n; i++ {
		tmpNext := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmpNext
	}
	/* 执行之后的结果
					 prem			pre  next
	                  |				 |	  |
		dummy -> 1 -> 2 -> 3 <- 4 <- 5    6 -> 7
										  |
										 cur
	*/
	// 缝合链表
	prem.Next.Next = cur //先缝合后半部分
	prem.Next = pre
	return dummy.Next
}

// 分组反转链表
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

func NewList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	tail := &ListNode{}
	head := tail
	for _, i := range arr {
		tail.Next = &ListNode{i, nil}
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
