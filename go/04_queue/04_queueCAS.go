package queue

import (
	"fmt"
	"sync/atomic"
)

// 实现无锁的CAS并发队列
// TODO 版本更新与指针移动是否要放在一个事务里？

// CASQueue 基于环形队列实现
type CASQueue struct {
	q         []interface{}
	head      int
	tail      int
	capacity  int
	enVersion int64
	deVersion int64
}

// NewCASQueue 新建环形队列
func NewCASQueue(cap int) *CASQueue {
	return &CASQueue{
		q:         make([]interface{}, cap),
		head:      0,
		tail:      0,
		capacity:  cap,
		enVersion: 0,
	}
}

// EnQueue 入队
func (q *CASQueue) EnQueue(v interface{}) bool {
	// 对列满了
	if (q.tail+1)%q.capacity == q.head {
		return false
	}
	old := q.enVersion
	ok := atomic.CompareAndSwapInt64(&q.enVersion, old, old+1)
	if ok {
		q.q[q.tail] = v
		q.tail = (q.tail + 1) % q.capacity
		return true
	}
	return false
}

// DeCASQueue 出队
func (q *CASQueue) DeCASQueue() interface{} {
	// 空队列
	if q.head == q.tail {
		return nil
	}
	old := q.deVersion
	ok := atomic.CompareAndSwapInt64(&q.deVersion, old, old+1)
	if ok {
		v := q.q[q.head]
		q.head = (q.head + 1) % q.capacity
		return v
	}
	return nil
}

// Print 打印
func (q *CASQueue) Print() {
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
