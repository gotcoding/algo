package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 给定一个链表，判断链表中是否有环。
// 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。
// 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
// 注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
// 如果链表中存在环，则返回 true 。 否则，返回 false 。

// 思路：
// 使用快慢指针，慢指针每次后移一个节点，快指针每次移动两个节点。
// 如果存在环，则快指针会追上慢指针

func hasCircle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// 以下为辅助测试代码

// ListNode 节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// NewListNode 新建节点
func NewListNode(v int) *ListNode {
	return &ListNode{
		Val:  v,
		Next: nil,
	}
}

// 生产随机链表
func randomList(maxLength int, maxValue int) *ListNode {
	rand.Seed(time.Now().Unix())
	length := rand.Intn(maxLength)
	head := NewListNode(0)
	pre := head
	for i := 0; i < length; i++ {
		v := rand.Intn(maxValue) - rand.Intn(maxValue) //正负随机
		newNode := NewListNode(v)
		pre.Next = newNode
		pre = newNode
	}
	return head.Next
}

// 获取尾节点
func getTail(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	for head.Next != nil {
		head = head.Next
	}
	return head
}

// 打印链表
func printList(head *ListNode) {
	if head == nil {
		fmt.Println("empty list")
	}
	for head != nil {
		fmt.Printf("%v -> ", head.Val)
		head = head.Next
	}
	fmt.Printf("\n")
}

func main() {
	ll := randomList(16, 100)
	printList(ll)
	tail := getTail(ll)
	fmt.Println(tail.Val)
	fmt.Println(hasCircle(ll))
	// 构造环
	tail.Next = ll
	fmt.Println(hasCircle(ll))
}
