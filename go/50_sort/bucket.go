package sort

// 桶排序的基本思路是将数据根据计算规则来分组，并将数据依次分配到对应分组中。
// 分配时可能出现某分组里有多个数据，那么再进行分组内排序。
// 然后得到了有序分组，最后把它们依次取出来放到数组中即实现了整体排序。

func getMax(arr []int) int {
	m := arr[0]
	for _, v := range arr {
		if m < v {
			m = v
		}
	}
	return m
}

// bucketSort
func bucketSort(arr []int) []int {
	l := len(arr)
	max := getMax(arr)
	buckets := make([][]int, l)

	// 分配入桶
	index := 0
	for i := 0; i < l; i++ {
		index = arr[i] * (l - 1) / max
		buckets[index] = append(buckets[index], arr[i])
	}

	// 桶内排序
	tmPos := 0
	for i := 0; i < l; i++ {
		l := len(buckets[i])
		if l > 0 {
			quickSort(buckets[i])
			copy(arr[tmPos:], buckets[i])
			tmPos += l
		}
	}
	return arr
}
