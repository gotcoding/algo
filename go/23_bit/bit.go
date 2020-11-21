package bit

// Swap 交换两个数
func Swap(a, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}

// 一个数组中有一种数出现了奇数次，其他数都出现了偶数次，怎么找到并打印这种数
func findEvenTimes(arr []int) int {
	eor := 0
	for _, v := range arr {
		eor ^= v
	}
	return eor
}

// arr中，有两种数，出现奇数次，找到并打印出来
func findEvenTimesDouble(arr []int) (int, int) {
	eor := 0
	for _, v := range arr {
		eor ^= v
	}
	rightOne := eor & (-eor) // 提取出最右的1
	onlyOne := 0
	for _, v := range arr {
		//  arr[1] =  111100011110000
		// rightOne=  000000000010000
		if (v & rightOne) != 0 {
			onlyOne ^= v
		}
	}
	return onlyOne, eor ^ onlyOne
}
