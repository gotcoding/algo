package sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	s := []int{4, 2, 9, 1, 8, 3, 5}
	fmt.Println(mergeSort(s))
}
