package stack

import (
	"fmt"
	"testing"
)

func TestLinkedListPush(t *testing.T) {
	s := NewLinkedListStack()
	for i := 0; i < 5; i++ {
		s.Push(i)
	}
	s.Print()
}

func TestLinkedListPop(t *testing.T) {
	s := NewLinkedListStack()
	for i := 0; i < 5; i++ {
		s.Push(i)
	}
	for i := 0; i < 6; i++ {
		fmt.Println(s.Pop())
	}
}
