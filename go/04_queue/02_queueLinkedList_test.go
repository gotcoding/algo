package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewLinkedListQueue()
	q.EnQueue("Tom")
	q.EnQueue("Jack")
	q.EnQueue("Mongo")
	q.Print()
	fmt.Println(q.DeQueue())
	q.Print()
}
