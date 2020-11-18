package sort_dev

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestBuild(t *testing.T) {
	l := 20
	arr := make([]int, l)
	for i := 0; i < l; i++ {
		arr[i] = int(rand.Intn(1000))
	}
	fmt.Println(arr)
	heapSort(arr)
	fmt.Println(arr)
}
