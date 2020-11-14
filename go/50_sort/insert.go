package sort

// 插入排序
// 对于未排序的数据，在已排序的序列中从后向前扫描，有点像打扑克牌时理牌的动作
// 时间复杂度为O(n^2)，最好为O(n)，空间复杂度为O(1)
func insertSort(d []int) []int {
	for i := 1; i < len(d); i++ {
		tmp := d[i]
		for j := i - 1; j >= 0; j-- {
			if tmp < d[j] {
				// 如果插入的元素小，交换位置。往后挪一位
				d[j+1] = d[j]
				// 将前面的数设置为当前需要交换的数
				d[j] = tmp
			} else { // 由于是已经排序好的，则不需要再次比较。
				break
			}
		}
	}
	return d
}
