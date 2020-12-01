package main

import (
	"fmt"
	"time"
)

// 链接：https://leetcode-cn.com/problems/edit-distance
// 给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。

// 你可以对一个单词进行如下三种操作：
// 插入一个字符
// 删除一个字符
// 替换一个字符

// 示例 1：
// 输入：word1 = "horse", word2 = "ros"
// 输出：3
// 解释：
// horse -> rorse (将 'h' 替换为 'r')
// rorse -> rose (删除 'r')
// rose -> ros (删除 'e')

func main() {
	word1 := "intention"
	word2 := "execution"
	start := time.Now()
	fmt.Println(minDistance(word1, word2))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println(dp2(word1, word2))
	fmt.Println(time.Since(start))
}

var _s1, _s2 string
var m [][]int

func minDistance(word1 string, word2 string) int {
	_s1, _s2 = word1, word2
	m = make([][]int, len(word1)+1)
	for i := 0; i <= len(word1); i++ {
		m[i] = make([]int, len(word2)+1)
	}
	return dp(len(word1)-1, len(word2)-1)
}

func dp(i, j int) int {
	// base case
	if i == -1 {
		return j + 1
	}
	if j == -1 {
		return i + 1
	}
	if m[i][j] > 0 {
		return m[i][j]
	}
	if _s1[i] == _s2[j] {
		m[i][j] = dp(i-1, j-1)

	} else {
		m[i][j] = min(
			dp(i, j-1)+1,   //插入
			dp(i-1, j)+1,   //删除
			dp(i-1, j-1)+1, //替换
		)
	}
	return m[i][j]
}

func min(args ...int) int {
	min := args[0]
	for _, i := range args {
		if min > i {
			min = i
		}
	}
	return min
}

func dp2(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	// 空字符串为0
	dp[0][0] = 0
	// dp[i][0] 表示编辑空字符串的代价，也就是将所有的字符删除的代价
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	// dp[0][i] 表示插入字符串的代价
	for i := 1; i <= n; i++ {
		dp[0][i] = i
	}
	// 自底向上求解
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(
					dp[i-1][j]+1,
					dp[i][j-1]+1,
					dp[i-1][j-1]+1)
			}
		}
	}
	return dp[m][n]
}
