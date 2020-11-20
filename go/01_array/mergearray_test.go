package arraydev

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMergeArray(t *testing.T) {
	arr1 := getSortedRandomArray(20)
	arr2 := getSortedRandomArray(10)
	fmt.Println(arr1)
	fmt.Println(arr2)
	arr := mergeArray(arr1, arr2)
	fmt.Println(arr)
}

func getSortedRandomArray(length int) []int {
	arr := make([]int, length)
	pre := 0
	for i := 0; i < length; i++ {
		tmp := rand.Intn(100)
		arr[i] = pre + tmp
		pre = tmp + pre
	}
	return arr
}
