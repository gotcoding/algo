package queue

import "fmt"

/*
 使用数组构建循环队列
*/

// CircularQueue 环形队列
type CircularQueue struct {
	q        []interface{}
	head     int
	tail     int
	capacity int
}

// NewCircularQueue 新建环形队列
func NewCircularQueue(cap int) *CircularQueue {
	return &CircularQueue{
		q:        make([]interface{}, cap),
		head:     0,
		tail:     0,
		capacity: cap,
	}
}

// EnQueue 入队
func (q *CircularQueue) EnQueue(v interface{}) bool {
	// 对列满了
	if (q.tail+1)%q.capacity == q.head {
		return false
	}
	q.q[q.tail] = v
	q.tail = (q.tail + 1) % q.capacity
	return true
}

// DeQueue 出队
func (q *CircularQueue) DeQueue() interface{} {
	// 空队列
	if q.head == q.tail {
		return nil
	}
	v := q.q[q.head]
	q.head = (q.head + 1) % q.capacity
	return v
}

// Print 打印
func (q *CircularQueue) Print() {
	if q.head == q.tail {
		fmt.Println("empty queue")
	}
	result := "head"
	i := q.head
	for true {
		result += fmt.Sprintf("<-%+v", q.q[i])
		i = (i + 1) % q.capacity
		if i == q.tail {
			break
		}
	}
	result += "<-tail"
	fmt.Println(result)
}
