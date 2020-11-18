package sort_dev

// 希尔排序
func shellSort(d []int) []int {
	for step := len(d) / 2; step > 0; step = step / 2 {
		for i := step; i < len(d); i++ {
			tmp := d[i]
			for j := i - step; j >= 0; j -= step {
				if tmp < d[j] {
					d[j+step] = d[j]
					d[j] = tmp
					tmp = d[j]
				} else {
					break
				}
			}
			// fmt.Printf("%d", step)
			// fmt.Println(d)
		}
	}
	return d
}
