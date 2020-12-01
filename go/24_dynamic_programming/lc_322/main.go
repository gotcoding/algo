package main

import (
	"fmt"
	"math"
)

// 链接：https://leetcode-cn.com/problems/coin-change
// 给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。
// 如果没有任何一种硬币组合能组成总金额，返回 -1。
// 你可以认为每种硬币的数量是无限的。

// 思路
// 1.对硬币面值排序
// 2.用期望面值按从大到小逐个去除

func main() {
	// fmt.Println(coinChange([]int{1, 2, 5}, 11))
	// fmt.Println(coinChange([]int{2}, 3))
	// fmt.Println(coinChange([]int{1}, 0))
	fmt.Println(coinChange([]int{186, 419, 83, 408}, 6249))
	fmt.Println(dp2([]int{186, 419, 83, 408}, 6249))
}

func coinChange(coins []int, amount int) int {
	return dp(coins, amount)
}

var m map[int]int = make(map[int]int)

// 动态规划，自上而下，逐个扣减，递归求解
// 1.问题可以拆分，子问题与父问题的模式一样，都是求最小硬币个数，只是amount值不同；
// 2.终止条件是amount为0时，0个硬币
// 3.模式是父的最小硬币数等于多个子最小硬币数加1
func dp(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	if v, ok := m[amount]; ok {
		return v
	}
	// 求最小值，初始化为最大值
	res := math.MaxInt32
	for _, coin := range coins {
		subProblem := dp(coins, amount-coin)
		// 子问题无解，跳过
		if subProblem == -1 {
			continue
		}
		res = min(res, 1+subProblem)
	}
	tmp := 0
	if res != math.MaxInt32 {
		tmp = res
	} else {
		tmp = -1
	}
	m[amount] = tmp
	return tmp
}

// 动态规划，迭代求解
func dp2(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = -1
		for _, c := range coins {
			if c > i || dp[i-c] == -1 {
				continue
			}
			count := dp[i-c] + 1
			if dp[i] == -1 || count < dp[i] {
				dp[i] = count
			}
		}
	}
	return dp[amount]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
