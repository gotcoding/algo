package queue

import (
	"fmt"
	"testing"
)

func TestCircularQueue(t *testing.T) {
	q := NewCircularQueue(6)
	for i := 0; i < 6; i++ {
		q.EnQueue(i)
	}
	q.Print()
	fmt.Println(q.DeQueue(), q.head)
	q.Print()
	q.EnQueue(100)
	q.Print()
	q.DeQueue()
	q.EnQueue(300)
	q.Print()
}

func TestCQ(t *testing.T) {
	q := NewCQ(6)
	for i := 0; i < 6; i++ {
		q.EnQueue(i)
	}
	q.Print()
	q.DeQueue()
	q.Print()
	q.EnQueue(10)
	fmt.Println(q.EnQueue(100))
	q.Print()
	fmt.Println(q.DeQueue())
}
