package queue

import (
	"errors"
	"fmt"
)

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

// CQ 循环队列 使用size简化
type CQ struct {
	head int
	tail int
	cap  int
	size int
	data []int
}

// NewCQ 新建环形队列
func NewCQ(cap int) *CQ {
	return &CQ{
		head: 0,
		tail: 0,
		cap:  cap,
		size: 0,
		data: make([]int, cap),
	}
}

// EnQueue 入队
func (q *CQ) EnQueue(v int) error {
	if q.size == q.cap {
		return errors.New("full queue")
	}
	q.size++
	q.data[q.tail] = v
	q.tail = (q.tail + 1) % q.cap
	return nil
}

// DeQueue 出队
func (q *CQ) DeQueue() (int, error) {
	if q.size == 0 {
		return 0, errors.New("empty queue")
	}
	q.size--
	d := q.data[q.head]
	q.head = (q.head + 1) % q.cap
	return d, nil
}

// Print 打印
func (q *CQ) Print() {
	if q.size == 0 {
		fmt.Println("empty queue")
	}
	result := "head"
	i := q.head
	for true {
		result += fmt.Sprintf("<-%+v", q.data[i])
		i = (i + 1) % q.cap
		if i == q.tail {
			break
		}
	}
	result += "<-tail"
	fmt.Println(result)
}
