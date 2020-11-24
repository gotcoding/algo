package main

import "fmt"

// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

// 示例:
// 给定 nums = [2, 7, 11, 15], target = 9
// 因为 nums[0] + nums[1] = 2 + 7 = 9
// 所以返回 [0, 1]

// 思路
// 将数组转换到map里，采用value:index的格式
// 然后再查询map，找合适的值，时间复杂度为O(n)，空间复杂度为O(n)

func main() {
	arr := []int{2, 7, 11, 15}
	fmt.Println(twoSum(arr, 9))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, 0)
	for i, v := range nums {
		if p, ok := m[target-v]; ok {
			return []int{i, p}
		}
		m[v] = i
	}
	return nil
}
