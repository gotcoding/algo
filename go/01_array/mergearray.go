package arraydev

// MergeSortedArrays 合并多个有序数组
// 思路：使用归并的方式来合并
func MergeSortedArrays(arrs [][]int) []int {
	return merge(arrs, 0, len(arrs)-1)
}

// 使用归并的方式来合并
func merge(arrs [][]int, left, right int) []int {
	if left == right {
		return arrs[left]
	}
	if left > right {
		return nil
	}
	mid := left + (right-left)>>2
	return MergeTwoArray(merge(arrs, left, mid), merge(arrs, mid+1, right))
}

// MergeTwoArray 合并两个有序数组
func MergeTwoArray(arr1, arr2 []int) []int {
	if len(arr1) == 0 {
		return arr2
	}
	if len(arr2) == 0 {
		return arr1
	}
	res := make([]int, 0)
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] > arr2[j] {
			res = append(res, arr2[j])
			j++
		} else {
			res = append(res, arr1[i])
			i++
		}
	}
	if i == len(arr1) {
		res = append(res, arr2[j:]...)
	}
	if j == len(arr2) {
		res = append(res, arr1[i:]...)
	}
	return res
}
