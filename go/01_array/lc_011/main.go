package main

import "fmt"

// 给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
// 说明：你不能倾斜容器。

// 思路1
// 两次for循环遍历，记录最大值
// 时间复杂度为O(n^2)，空间复杂度O(1)
func maxArea(height []int) int {
	max := 0
	for i := len(height) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			v := min(height[i], height[j]) * (i - j)
			if v > max {
				max = v
			}
		}
	}
	return max
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 思路2：双指针
// 时间复杂度为O(n)，空间复杂度O(1)
func maxArea1(height []int) int {
	L := 0
	R := len(height) - 1
	max := 0
	for L < R {
		v := min(height[R], height[L]) * (R - L)
		if v > max {
			max = v
		}
		if height[L] > height[R] {
			R--
		} else {
			L++
		}
	}
	return max
}

func main() {
	arr1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea1(arr1) == 49)
	arr2 := []int{1, 1}
	fmt.Println(maxArea1(arr2) == 1)
	arr3 := []int{4, 3, 2, 1, 4}
	fmt.Println(maxArea1(arr3) == 16)
	arr4 := []int{1, 2, 1}
	fmt.Println(maxArea1(arr4) == 2)
}
