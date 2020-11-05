package skiplist

import (
	"fmt"
	"testing"
)

func TestNewSkipList(t *testing.T) {
	sl := NewSkipList(5)
	sl.Insert(95, "Tom")
	sl.Insert(100, "Jason")
	sl.Insert(78, "Lily")
	sl.Insert(60, "Lucy")
	sl.Insert(102, "Jessca")
	// sl.Delete(100)
	sl.Print()
	fmt.Println("78:", sl.Search(78).data)
}

func TestRandomLevel(t *testing.T) {
	sl := NewSkipList(10)
	for i := 0; i < 10000; i++ {
		r := sl.randomLevel()
		if r > 3 {
			fmt.Println(r)
		}
	}
}
