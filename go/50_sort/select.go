package sort

// 选择排序
// 每次选择一个最大或最小值放在已排序的序列末尾
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
