package search

// 二分查找
func binarySearch(arr []int, v int) bool {
	if arr == nil {
		return false
	}
	L := 0
	R := len(arr) - 1
	for L < R {
		mid := L + ((R - L) >> 1)
		if arr[mid] == v {
			return true
		} else if arr[mid] > v {
			R = mid - 1
		} else {
			L = mid + 1
		}
	}
	return arr[L] == v
}
