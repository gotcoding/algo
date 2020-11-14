package sort

import (
	"fmt"
	"testing"
)

func TestSelectSort(t *testing.T) {
	d := []int{4, 2, 9, 1, 8, 3, 5}
	res := selectSort(d)
	fmt.Println(res)
}
