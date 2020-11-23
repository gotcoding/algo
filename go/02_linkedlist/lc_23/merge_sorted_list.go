package lc23

// 给你一个链表数组，每个链表都已经按升序排列。
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。

// 示例 1：
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
// 输出：[1,1,2,3,4,4,5,6]

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewNode(v int) *ListNode {
	return &ListNode{
		Val:  v,
		Next: nil,
	}
}

func NewList(arr []int) *ListNode {
	head := NewNode(0)
	tail := head
	if len(arr) == 0 {
		return nil
	}
	for _, v := range arr {
		node := NewNode(v)
		tail.Next = node
		tail = tail.Next
	}
	return head.Next
}

// 思路1：顺序合并，逐个合并
func mergeKLists(lists []*ListNode) *ListNode {
	var res *ListNode
	for _, l := range lists {
		res = mergeTwoList(res, l)
	}
	return res
}

// 思路2:分治合并
func mergeKLists2(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	if l > r {
		return nil
	}
	mid := l + (r-l)>>2
	return mergeTwoList(merge(lists, l, mid), merge(lists, mid+1, r))
}

func mergeTwoList(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	a, b := list1, list2
	head := &ListNode{}
	tail := head
	for a != nil && b != nil {
		if a.Val > b.Val {
			tail.Next = b
			b = b.Next
		} else {
			tail.Next = a
			a = a.Next
		}
		tail = tail.Next
	}
	if a == nil {
		tail.Next = b
	}
	if b == nil {
		tail.Next = a
	}
	return head.Next
}
