package stack

import (
	"fmt"
	"testing"
)

func TestPush(t *testing.T) {
	s := NewArrayStack()
	for i := 0; i < 5; i++ {
		s.Push(i)
	}
	s.Print()
}

func TestPop(t *testing.T) {
	s := NewArrayStack()
	for i := 0; i < 5; i++ {
		s.Push(i)
	}
	count := s.top
	for i := 0; i <= count; i++ {
		fmt.Println(s.Pop())
	}
}

func TestTop(t *testing.T) {
	s := NewArrayStack()
	for i := 0; i < 5; i++ {
		s.Push(i)
	}
	s.Print()
	fmt.Println(s.Top())
}
