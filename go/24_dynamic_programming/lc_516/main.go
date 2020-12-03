package main

import "fmt"

// 最长回文子序列
// 给定一个字符串 s ，找到其中最长的回文子序列，并返回该序列的长度。可以假设 s 的最大长度为 1000 。
// 示例 1:
// 输入: "bbbab"
// 输出: 4

// 解题思路：
// 状态
// f[i][j] 表示 s 的第 i 个字符到第 j 个字符组成的子串中，最长的回文序列长度是多少。

// 转移方程
// 如果 s 的第 i 个字符和第 j 个字符相同的话
// f[i][j] = f[i + 1][j - 1] + 2
// 如果 s 的第 i 个字符和第 j 个字符不同的话
// f[i][j] = max(f[i + 1][j], f[i][j - 1])
// 然后注意遍历顺序，i 从最后一个字符开始往前遍历，j 从 i + 1 开始往后遍历，这样可以保证每个子问题都已经算好了。

// 初始化
// f[i][i] = 1 单个字符的最长回文序列是 1

// 结果
// f[0][n - 1]

func main() {
	fmt.Println(longestPalindromeSubseq("bbbab"))
	fmt.Println(longestPalindromeSubseq("a"))
	fmt.Println(longestPalindromeSubseq("ac"))
	fmt.Println(longestPalindromeSubseq("acbc"))
	fmt.Println(longestPalindromeSubseq2("bbbab"))
	fmt.Println(longestPalindromeSubseq2("a"))
	fmt.Println(longestPalindromeSubseq2("ac"))
	fmt.Println(longestPalindromeSubseq2("acbc"))
}

func longestPalindromeSubseq(s string) int {
	l := len(s)
	mem := make([][]int, l)
	for i := 0; i < l; i++ {
		mem[i] = make([]int, l)
	}
	// base case
	for i := 0; i < l; i++ {
		mem[i][i] = 1
	}
	for i := l - 1; i >= 0; i-- {
		for j := i + 1; j < l; j++ {
			// 状态转移方程
			if s[i] == s[j] {
				mem[i][j] = mem[i+1][j-1] + 2
			} else {
				mem[i][j] = Max(mem[i+1][j], mem[i][j-1])
			}
		}
	}
	// 返回整个s的最长回文
	return mem[0][l-1]
}

func longestPalindromeSubseq2(s string) int {
	if len(s) == 1 {
		return 1
	}
	return dp(s, 0, len(s)-1)
}

// 动态规划，递归
func dp(s string, b, e int) int {
	// base case
	if e-b == 0 {
		return 1
	}
	if e-b < 0 {
		return 0
	}
	if s[b] == s[e] {
		return dp(s, b+1, e-1) + 2
	} else {
		return Max(dp(s, b+1, e), dp(s, b, e-1))
	}
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
