package main

import (
	"fmt"
	"strconv"
)

// 根据 逆波兰表示法，求表达式的值。
// 有效的运算符包括 +, -, *, / 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。

// 示例 1：
// 输入: ["2", "1", "+", "3", "*"]
// 输出: 9
// 解释: 该算式转化为常见的中缀算术表达式为：((2 + 1) * 3) = 9

func evalRPN(tokens []string) int {
	stack := make([]int64, 0)
	for _, str := range tokens {
		l := len(stack)
		switch str {
		case "+":
			stack[l-2] = stack[l-1] + stack[l-2]
			stack = stack[:l-1]
		case "-":
			stack[l-2] = stack[l-2] - stack[l-1]
			stack = stack[:l-1]
		case "*":
			stack[l-2] = stack[l-1] * stack[l-2]
			stack = stack[:l-1]
		case "/":
			stack[l-2] = stack[l-2] / stack[l-1]
			stack = stack[:l-1]
		default:
			v, _ := strconv.ParseInt(str, 10, 64)
			stack = append(stack, v)
		}
	}
	return int(stack[0])
}

func main() {
	tokens1 := []string{"2", "1", "+", "3", "*"}
	fmt.Println(evalRPN(tokens1))
	tokens3 := []string{"4", "13", "5", "/", "+"}
	fmt.Println(evalRPN(tokens3))
	tokens2 := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(tokens2))
}
