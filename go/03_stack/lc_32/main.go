package main

import "fmt"

// 给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。
// 示例 1:
// 输入: "(()"
// 输出: 2
// 解释: 最长有效括号子串为 "()"

// 思路
// 使用栈
// https://leetcode-cn.com/problems/longest-valid-parentheses/solution/shou-hua-tu-jie-zhan-de-xiang-xi-si-lu-by-hyj8/
func longestValidParentheses(s string) int {
	maxLength := 0
	stack := []int{}
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				maxLength = max(maxLength, i-stack[len(stack)-1])
			}
		}
	}
	return maxLength
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	str := "(()()(())"
	fmt.Println(longestValidParentheses(str))
}
