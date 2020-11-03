package queue

import (
	"fmt"
	"sync"
	"testing"
)

func TestCASEnQueue(t *testing.T) {
	var wg sync.WaitGroup
	q := NewCASQueue(5)
	threadNum := 500
	wg.Add(threadNum)
	for i := 0; i < threadNum; i++ {
		go func(i int) {
			defer wg.Done()
			spinNum := 0
			for {
				ok := q.EnQueue(1)
				if ok {
					break
				} else {
					spinNum++
				}
			}
			fmt.Printf("thread,%d,spinnum,%d\n", i, spinNum)
		}(i)
	}
	go func() {
		total := 0
		for {
			v, ok := q.DeCASQueue().(int)
			if ok {
				total += v
			}
			if q.deVersion == int64(threadNum) {
				break
			}
		}
		fmt.Println(total)
	}()
	wg.Wait()
	fmt.Println(q.enVersion, q.deVersion)
}
