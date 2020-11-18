package sort_dev

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	d := []int{4, 2, 9, 1, 8, 3, 5}
	quickSort(d)
	fmt.Println(d)
}
