package sort

import (
	"fmt"
	"testing"
)

func TestCountingSort(t *testing.T) {
	arr := []int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1}
	fmt.Println(arr)
	countingSort(arr)
	fmt.Println(arr)
}
