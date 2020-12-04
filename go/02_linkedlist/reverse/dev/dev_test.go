package dev

import (
	"fmt"
	"testing"
)

func TestNewList(t *testing.T) {
	list := NewList([]int{1, 2, 3, 4, 5, 6, 7})
	Print(list)
}

func TestHasCircle(t *testing.T) {
	list := NewList([]int{1, 2, 3, 4, 5, 6, 7})
	Print(list)
	fmt.Println(hasCircle(list))
	list.Next.Next.Next.Next.Next = list
	fmt.Println(hasCircle(list))
}

func TestMergeList(t *testing.T) {
	l1 := NewList([]int{1, 2, 3, 4, 5, 6, 7})
	l2 := NewList([]int{4, 5})
	l3 := NewList([]int{})
	Print(mergeList(l1, l1))
	Print(mergeList(l1, l2))
	Print(mergeList(l3, l1))
}

func TestReverse(t *testing.T) {
	l1 := NewList([]int{1, 2, 3, 4, 5, 6, 7})
	l2 := NewList([]int{1})
	Print(reverse(l1))
	Print(reverse(l2))
}

func TestReverseBetween(t *testing.T) {
	l1 := NewList([]int{1, 2, 3, 4, 5, 6, 7})
	// Print(reverseBetween(l1, 2, 4))
	Print(reverseBetween(l1, 1, 4))
}

func TestReverseGroupK(t *testing.T) {
	l1 := NewList([]int{1, 2, 3, 4, 5, 6, 7})
	// head, _ := reverseNodes(l1, l1.Next.Next.Next)
	// Print(head)
	Print(reverseGroupK(l1, 4))
}
