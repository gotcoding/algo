package bit

import (
	"fmt"
	"testing"
)

func TestSwap(t *testing.T) {
	a := 3
	b := 7
	a, b = Swap(a, b)
	fmt.Println(a, b)
}

func TestFindEvenTimes(t *testing.T) {
	arr := []int{4, 2, 3, 2, 3, 3, 4, 5, 3, 2, 4, 2, 5}
	fmt.Println(findEvenTimes(arr))
}

func TestFindEvenTimesDouble(t *testing.T) {
	arr := []int{4, 2, 3, 2, 3, 3, 4, 5, 3, 2, 4, 5}
	fmt.Println(findEvenTimesDouble(arr))
}
