package sort_dev

import "math/rand"

// 随机数组
func generateRandomArray(maxSize, maxValue int) []int {
	arr := make([]int, rand.Intn(maxSize))
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(maxValue)
	}
	return arr
}

// 数组比对
func isEquel(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i, v := range arr1 {
		if arr2[i] != v {
			return false
		}
	}
	return true
}
