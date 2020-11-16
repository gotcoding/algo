package sort

// 选择排序
// 分已排序区间和未排序区间
// 每次会从未排序区间中找到最小的元素，将其放到已排序区间的末尾。
// 时间复杂度为O(n^2)，最优为O(n^2)，空间复杂度为O(1)
func selectSort(d []int) []int {
	for i := 0; i < len(d); i++ {
		for j := i; j < len(d); j++ {
			if d[j] < d[i] {
				d[i], d[j] = d[j], d[i]
			}
		}
	}
	return d
}
