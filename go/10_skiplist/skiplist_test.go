package skiplist

import (
	"testing"
)

func TestSkipList(t *testing.T) {
	sl := NewSkipList()
	sl.Insert("tom", 95)
	sl.Insert("json", 88)
	sl.Insert("jack", 100)
	sl.Insert("cat", 60)
	t.Log(sl.head.forwards[0])
	t.Log(sl.head.forwards[0].forwards[0])
	t.Log(sl.head.forwards[0].forwards[0].forwards[0])
	t.Log(sl.head.forwards[0].forwards[0].forwards[0].forwards[0])
	t.Log(sl)
	t.Log(sl.Find("json", 88))
}
