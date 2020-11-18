package heaptest

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestBuild(t *testing.T) {
	h := NewHeap(20)
	for i := 0; i < h.Length; i++ {
		h.Value[i] = int(rand.Intn(1000))
	}
	fmt.Println(h.Value)
	h.build()
}
