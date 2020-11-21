package sort_dev

import "math/rand"

// 随机数组
func generateRandomArray(maxSize, maxValue int) []int {
	arr := make([]int, rand.Intn(maxSize))
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(maxValue+1) - rand.Intn(maxValue)
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

// 生成一个随机有序数组
func generateSoredArray(length uint) []int {
	arr := make([]int, length)
	var tmp int = 0
	for i := uint(0); i < length; i++ {
		tmp += rand.Intn(100)
		arr[i] = tmp
	}
	return arr
}
