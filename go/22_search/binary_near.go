package search

// 在一个有序数组中，找>=某个数最左侧的位置
func bsNearLeft(arr []int, v int) int {
	L, R := 0, len(arr)-1
	idx := -1
	for L <= R {
		mid := L + (R-L)>>1
		if arr[mid] >= v {
			idx = mid
			R = mid - 1
		} else {
			L = mid + 1
		}
	}
	return idx
}

// 在一个有序数组中，找<=某个数最右侧的位置
func bsNearRight(arr []int, v int) int {
	L, R := 0, len(arr)-1
	idx := -1
	for L <= R {
		mid := L + (R-L)>>1
		if arr[mid] <= v {
			idx = mid
			L = mid + 1
		} else {
			R = mid - 1
		}
	}
	return idx
}
