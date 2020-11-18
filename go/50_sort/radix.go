package sort_dev

// radixSort 基数排序
func radixSort(arr []int) {
	max := getMax(arr)
	maxFigure := 1
	for max/10 > 0 {
		maxFigure++
		max = max / 10
	}
	figrueSort(arr, maxFigure)
}

func figrueSort(arr []int, mfg int) {
	num := 1
	for i := 0; i < mfg; i++ {
		var bucket [10][]int
		var result []int
		for _, s := range arr {
			n := s / num % 10
			bucket[n] = append(bucket[n], s)
		}

		// append
		for i := 0; i < 10; i++ {
			result = append(result, bucket[i]...)
		}

		for i := range result {
			arr[i] = result[i]
		}
		num *= 10
	}
}
