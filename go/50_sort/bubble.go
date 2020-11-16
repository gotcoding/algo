package sort

// 冒泡排序
// 每次比较相邻的两个数，每一轮冒泡出一个最大或者最小值放在已排序的最前面。
// 时间复杂度为O(n^2)，空间复杂度O(1)
func bubbleSort(d []int) []int {
	for i := 0; i < len(d); i++ {
		for j := 0; j < len(d)-i-1; j++ {
			if d[j] > d[j+1] {
				d[j], d[j+1] = d[j+1], d[j]
			}
		}
	}
	return d
}

// 优化：记录每轮的交换次数，如果数组是有序的，则跳出循环
// 时间复杂度为O(n^2)，最好则为O(n)，平均是O(n^2)
func bubbleSortPro(d []int) []int {
	for i := 0; i < len(d); i++ {
		count := 0
		for j := 0; j < len(d)-i-1; j++ {
			if d[j] > d[j+1] {
				d[j], d[j+1] = d[j+1], d[j]
				count++
			}
		}
		// fmt.Println(count)
		if count == 0 {
			break
		}
	}
	return d
}
