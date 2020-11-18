package sort_dev

// 选择排序
// 分已排序区间和未排序区间
// 每次会从未排序区间中找到最小的元素，将其放到已排序区间的末尾。
// 时间复杂度为O(n^2)，最优为O(n^2)，空间复杂度为O(1)
func selectSort(arr []int) {
	l := len(arr)
	if arr == nil || l < 2 {
		return
	}

	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
