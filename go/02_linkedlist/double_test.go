package linkedlist

import (
	"testing"
)

func TestInsert(t *testing.T) {
	dl := NewDoubleLinked()
	for i := 0; i < 10; i++ {
		dl.InsertToTail(i)
	}
	dl.Print()
}
func TestDoubleReverse(t *testing.T) {
	dl := NewDoubleLinked()
	for i := 0; i < 10; i++ {
		dl.InsertToTail(i)
	}
	dl.Print()
	dl.Reverse()
	dl.Print()
}
