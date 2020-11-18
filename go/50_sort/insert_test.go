package sort_dev

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	d := []int{2, 9, 5, 3, 6, 1, 8}
	fmt.Println(insertSort(d))
}
