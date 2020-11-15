package sort

import (
	"fmt"
	"testing"
)

func TestShellSort(t *testing.T) {
	d := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(shellSort(d))
}
