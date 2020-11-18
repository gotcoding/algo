package sort

// 插入排序
// 1.将数组中的数据分为两个区间，已排序区间和未排序区间
// 2.初始已排序区间只有一个元素，就是数组的第一个元素
// 3.取未排序区间中的元素，在已排序区间中找到合适的插入位置将其插入，并保证已排序区间数据一直有序
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