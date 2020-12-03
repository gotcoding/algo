package main

import "fmt"

// 两个字符串的删除操作 链接：https://leetcode-cn.com/problems/delete-operation-for-two-strings
// 给定两个单词 word1 和 word2，找到使得 word1 和 word2 相同所需的最小步数，每步可以删除任意一个字符串中的一个字符。

// 示例：
// 输入: "sea", "eat"
// 输出: 2
// 解释: 第一步将"sea"变为"ea"，第二步将"eat"变为"ea"

func main() {
	fmt.Println(minDistance("sea", "eat"))
	fmt.Println(minDistance2("sea", "eat"))
}

var mem [][]int

func minDistance(word1 string, word2 string) int {
	// return dp(word1, word2, 0, 0)
	mem = make([][]int, len(word1)+1)
	for i := 0; i < len(word1)+1; i++ {
		mem[i] = make([]int, len(word2)+1)
	}
	return dp1(word1, word2, 0, 0)
}

// 递归暴力破解
func dp(word1, word2 string, m, n int) int {
	// base case
	if m == len(word1) {
		return len(word2) - n
	}
	if n == len(word2) {
		return len(word1) - m
	}
	if word1[m] == word2[n] {
		return dp(word1, word2, m+1, n+1)
	} else {
		return Min(dp(word1, word2, m+1, n), dp(word1, word2, m, n+1)) + 1
	}
}

// 递归暴力破解优化
func dp1(word1, word2 string, m, n int) int {
	// base case
	if m == len(word1) {
		return len(word2) - n
	}
	if n == len(word2) {
		return len(word1) - m
	}
	if mem[m][n] > 0 {
		return mem[m][n]
	}
	if word1[m] == word2[n] {
		mem[m][n] = dp(word1, word2, m+1, n+1)
	} else {
		mem[m][n] = Min(dp(word1, word2, m+1, n), dp(word1, word2, m, n+1)) + 1
	}
	return mem[m][n]
}

// 先求最长自序列
func minDistance2(word1 string, word2 string) int {
	max := dp3(word1, word2)
	return len(word1) - max + len(word2) - max
}
func dp3(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	mem := make([]int, n+1)
	for i := 1; i <= m; i++ {
		last := 0 //记录左上方的值
		for j := 1; j <= n; j++ {
			tmp := mem[j] //记录上方的值
			if text1[i-1] == text2[j-1] {
				mem[j] = last + 1
			} else {
				mem[j] = Max(tmp, mem[j-1])
			}
			last = tmp
		}
	}
	return mem[n]
}

func Max(args ...int) int {
	max := args[0]
	for _, i := range args {
		if i > max {
			max = i
		}
	}
	return max
}

func Min(args ...int) int {
	min := args[0]
	for _, i := range args {
		if i < min {
			min = i
		}
	}
	return min
}
