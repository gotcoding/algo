package sort_dev

import (
	"fmt"
	"testing"
)

func TestRadixSort(t *testing.T) {
	ce := []int{6, 4, 7, 3, 9, 5, 1, 13, 0, 8, 23, 10, 2}
	fmt.Printf("before: %d\n", ce)

	radixSort(ce)
	fmt.Printf("after：%d\n", ce)
}
