package recursion

// 有n阶楼梯，每次可以走1阶或者2阶，请问有多少种走法？
func timesRec(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return timesRec(n-1) + timesRec(n-2)
}

// Fibs 递归实现斐波拉契数列
type Fibs struct {
	dic map[int]int
	arr []int
}

// NewFibs 新建
func NewFibs(n int) *Fibs {
	return &Fibs{dic: make(map[int]int, n), arr: make([]int, n+1)}
}

func (f *Fibs) fibonacciDic(n int) int {
	if v, ok := f.dic[n]; ok {
		return v
	}
	if n == 1 {
		f.dic[n] = 1
		return 1
	}
	if n == 2 {
		f.dic[n] = 2
		return 2
	}
	res := timesRec(n-1) + timesRec(n-2)
	f.dic[n] = res
	return res
}

func (f *Fibs) fibonacciArr(n int) int {
	if f.arr[n] != 0 {
		return f.arr[n]
	}
	if n == 1 {
		f.arr[n] = 1
		return 1
	}
	if n == 2 {
		f.arr[n] = 2
		return 2
	}
	res := timesRec(n-1) + timesRec(n-2)
	f.arr[n] = res
	return res
}

func inFor(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	ret := 0
	pre := 2
	prepre := 1
	for i := 3; i <= n; i++ {
		ret = pre + prepre
		prepre = pre
		pre = ret
	}
	return ret
}
