package hashmap

import (
	"fmt"
	"strconv"
	"testing"
)

func TestHashMap(t *testing.T) {
	hm := NewHashMap(10)
	fmt.Println(hm.cap)
	for i := 0; i < 18; i++ {
		hm.Add(strconv.FormatInt(int64(i), 10), i)
	}
	hm.Print()
	fmt.Println(hm.len, hm.cap)
	fmt.Println(hm.Get("4"))
}
