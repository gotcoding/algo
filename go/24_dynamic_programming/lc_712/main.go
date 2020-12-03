package main

import "fmt"

//  两个字符串的最小ASCII删除和 链接：https://leetcode-cn.com/problems/minimum-ascii-delete-sum-for-two-strings
// 给定两个字符串s1, s2，找到使两个字符串相等所需删除字符的ASCII值的最小和。

// 示例 1:
// 输入: s1 = "sea", s2 = "eat"
// 输出: 231
// 解释: 在 "sea" 中删除 "s" 并将 "s" 的值(115)加入总和。
// 在 "eat" 中删除 "t" 并将 116 加入总和。
// 结束时，两个字符串相等，115 + 116 = 231 就是符合条件的最小和。

func main() {
	fmt.Println(minimumDeleteSum("sea", "eat"))
	fmt.Println(minimumDeleteSum("delete", "leet"))
}

func minimumDeleteSum(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	mem := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		mem[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				mem[i][j] = mem[i-1][j-1] + int(s1[i-1])
			} else {
				mem[i][j] = Max(mem[i-1][j], mem[i][j-1])
			}
		}
	}
	return asciiSum(s1) + asciiSum(s2) - 2*mem[m][n]
}

func asciiSum(s string) int {
	sum := 0
	for _, i := range s {
		sum += int(i)
	}
	return sum
}

func Max(args ...int) int {
	max := args[0]
	for _, i := range args {
		if max < i {
			max = i
		}
	}
	return max
}
