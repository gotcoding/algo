package sort_dev

import (
	"fmt"
	"sort"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	testCounts := 10
	maxSize := 100
	maxValue := 100
	success := true
	for i := 0; i < testCounts; i++ {
		arr1 := generateRandomArray(maxSize, maxValue)
		arr2 := make([]int, len(arr1))
		copy(arr2, arr1)
		// bubbleSort(arr1)
		bubbleSortPro(arr1)
		sort.Ints(arr2)
		if !isEquel(arr1, arr2) {
			fmt.Println(arr1)
			fmt.Println(arr2)
			success = false
			break
		}
	}
	if success {
		t.Log("Nice!")
	} else {
		t.Error("Fail!")
	}
}
