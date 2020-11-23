package lc23

import (
	"fmt"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	l1 := NewList([]int{1, 4, 5})
	l2 := NewList([]int{1, 3, 4})
	l3 := NewList([]int{2, 6})
	lists := []*ListNode{l1, l2, l3}
	res := mergeKLists(lists)
	for res != nil {
		fmt.Printf("->%v", res.Val)
		res = res.Next
	}
	fmt.Printf("\n")
}

func TestMergeKLists2(t *testing.T) {
	l1 := NewList([]int{1, 4, 5})
	l2 := NewList([]int{1, 3, 4})
	l3 := NewList([]int{2, 6})
	lists := []*ListNode{l1, l2, l3}
	res := mergeKLists2(lists)
	for res != nil {
		fmt.Printf("->%v", res.Val)
		res = res.Next
	}
	fmt.Printf("\n")
}
