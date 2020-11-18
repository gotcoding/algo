package sort_dev

func countingSort(arr []int) {
	l := len(arr)
	if l <= 1 {
		return
	}
	max := getMax(arr)
	c := make([]int, max+1)
	// 计算每个元素的个数，放入c中
	for i := 0; i < l; i++ {
		c[arr[i]]++
	}
	// 依次累加
	for i := 1; i <= max; i++ {
		c[i] += c[i-1]
	}
	// 临时数组res，存储排序之后的结果
	res := make([]int, l)
	for i := l - 1; i >= 0; i-- {
		idx := c[arr[i]] - 1
		res[idx] = arr[i]
		c[arr[i]]--
	}
	copy(arr, res)
}
