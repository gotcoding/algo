package queue

import "fmt"

// ArrayQueue 数组实现的队列
type ArrayQueue struct {
	q        []interface{}
	capacity int
	head     int
	tail     int
}

// NewArrayQueue 新建队列
func NewArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{make([]interface{}, n), n, 0, 0}
}

// EnQueue 入队
func (aq *ArrayQueue) EnQueue(v interface{}) bool {
	if aq.tail == aq.capacity {
		return false
	}
	aq.q[aq.tail] = v
	aq.tail++
	return true
}

// DeQueue 出队
func (aq *ArrayQueue) DeQueue() interface{} {
	if aq.head == aq.tail {
		return nil
	}
	v := aq.q[aq.head]
	aq.head++
	return v
}

// Print 打印
func (aq *ArrayQueue) Print() {
	if aq.head == aq.tail {
		fmt.Println("empty queue")
	}
	result := "head"
	for i := aq.head; i < aq.tail; i++ {
		result += fmt.Sprintf("<-%+v", aq.q[i])
	}
	result += "<-tail"
	fmt.Println(result)
}
