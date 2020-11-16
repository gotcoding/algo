package sort

// 归并排序
func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	result := merge(left, right)
	return result
}

func merge(left, right []int) []int {
	res := make([]int, 0)
	m, n := 0, 0
	l, r := len(left), len(right)
	for m < l && n < r {
		if left[m] > right[n] {
			res = append(res, right[n])
			n++
			continue
		}
		res = append(res, left[m])
		m++
	}
	res = append(res, left[m:]...)
	res = append(res, right[n:]...)
	return res
}
