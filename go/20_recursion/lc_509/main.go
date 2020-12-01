package main

import (
	"fmt"
	"time"
)

// 链接：https://leetcode-cn.com/problems/fibonacci-number
// 斐波那契数，通常用 F(n) 表示，形成的序列称为斐波那契数列。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：
// F(0) = 0,   F(1) = 1
// F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
// 给定 N，计算 F(N)。

func main() {
	start := time.Now()
	fmt.Println(fib1(20))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println(fib2(20))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println(fib3(20))
	fmt.Println(time.Since(start))
}

// 递归迭代
func fib1(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}
	return fib1(N-1) + fib1(N-2)
}

var m map[int]int = make(map[int]int)

// 使用MAP来缓存数据
func fib2(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 || N == 2 {
		return 1
	}
	if v, ok := m[N]; ok {
		return v
	} else {
		res := fib2(N-1) + fib2(N-2)
		m[N] = res
		return res
	}
}

// 使用迭代
func fib3(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 || N == 2 {
		return 1
	}
	pre, dPre := 1, 1
	res := 0
	for i := 3; i <= N; i++ {
		res = pre + dPre
		dPre = pre
		pre = res
	}
	return res
}
