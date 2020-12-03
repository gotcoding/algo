package main

import "fmt"

// 53. 最大子序和 https://leetcode-cn.com/problems/maximum-subarray
// 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

// 示例:
// 输入: [-2,1,-3,4,-1,2,1,-5,4]
// 输出: 6
// 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。

// 思路1：动态规划
// 状态：dp[i][j]代表nums中下标 i 到 j 之间最大连续数组之和

// 状态转移方程：
// 如果num[i]小于0，dp[i][j] = dp[i][j-1]
// 否则dp[i][j] = num[j] + dp[i+1][j]
// 遍历时 i 从后向前遍历，j 从 i+1 开始向后遍历，可以保证每个子结果都能算好

// 初始状态
// i=j 时dp[i][j] 为num[i]

// 结构
// dp[0][n-1]

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	// 初始状态
	for i := 0; i < n; i++ {
		dp[i][i] = nums[i]
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			// dp[i][j] = Max(dp[i][j-1], dp[i+1][j], dp[i+1][j-1])
			if nums[i] > 0 {
				dp[i][j] = nums[i] + dp[i][j-1]
			} else {
				dp[i][j] = nums[j] + dp[i+1][j]
			}
		}
	}
	return dp[0][n-1]
}

func Max(args ...int) int {
	max := 0
	for _, i := range args {
		if max < i {
			max = i
		}
	}
	return max
}
