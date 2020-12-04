package reverse

import "testing"

func TestNewList(t *testing.T) {
	list := NewList([]int{1, 2, 3, 4, 5, 6})
	Print(list)
}

func TestIteration(t *testing.T) {
	list := NewList([]int{1, 2, 3, 4, 5, 6})
	Print(list)
	newList := reverseIteration(list)
	Print(newList)
}

func TestReverseBetween(t *testing.T) {
	list := NewList([]int{1, 2, 3, 4, 5, 6})
	Print(list)
	r := reverseBetween(list, 1, 4)
	Print(r)
}

func TestReverseKGroup(t *testing.T) {
	list := NewList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	Print(list)
	l := reverseKGroup(list, 3)
	Print(l)
}
