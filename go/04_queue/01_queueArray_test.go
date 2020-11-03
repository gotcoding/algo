package queue

import (
	"fmt"
	"testing"
)

func TestEnQueue(t *testing.T) {
	aq := NewArrayQueue(100)
	aq.EnQueue("Tom")
	aq.EnQueue("Json")
	aq.Print()
}

func TestDeQueue(t *testing.T) {
	aq := NewArrayQueue(100)
	aq.EnQueue("Tom")
	aq.EnQueue("Json")
	aq.EnQueue("Cat")
	aq.Print()
	fmt.Println(aq.DeQueue())
	aq.Print()
}
