package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	s := []int{4, 2, 9, 1, 8, 3, 5}
	fmt.Println(bubbleSort(s))
	s1 := []int{4, 2}
	fmt.Println(bubbleSort(s1))
}

func TestBubbleSortPro(t *testing.T) {
	s := []int{4, 2, 9, 1, 8, 3, 5}
	fmt.Println(bubbleSortPro(s))
	// 最优
	s1 := []int{1, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(bubbleSortPro(s1))
	// 最差
	s2 := []int{9, 7, 5, 4, 3, 2, 1}
	fmt.Println(bubbleSortPro(s2))
}
