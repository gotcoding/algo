package arraydev

// 练习1: 合并两个有序数组
func mergeArray(arr1, arr2 []int) []int {
	if arr1 == nil {
		return arr2
	}
	if arr2 == nil {
		return arr1
	}
	arr := make([]int, 0)
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			arr = append(arr, arr1[i])
			i++
		} else {
			arr = append(arr, arr2[j])
			j++
		}
	}
	if i == len(arr1)-1 {
		arr = append(arr, arr2[j:]...)
	}
	if j == len(arr2)-1 {
		arr = append(arr, arr1[i:]...)
	}
	return arr
}
