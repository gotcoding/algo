package main

import (
	"fmt"
	"sort"
)

// 15
// 链接：https://leetcode-cn.com/problems/3sum
// 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？
// 请你找出所有满足条件且不重复的三元组。
// 注意：答案中不可以包含重复的三元组。

// 示例：
// 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
// 满足要求的三元组集合为：
// [
//   [-1, 0, 1],
//   [-1, -1, 2]
// ]

// 思路
// 标签：数组遍历，双指针
// 首先对数组进行排序，排序后固定一个数 nums[i]，再使用左右指针指向 nums[i]后面的两端，数字分别为 nums[L] 和 nums[R]，计算三个数的和 sum 判断是否满足为 0，满足则添加进结果集
// 如果 nums[i] 大于 0，则三数之和必然无法等于 0，结束循环
// 如果 nums[i] == nums[i-1]，则说明该数字重复，会导致结果重复，所以应该跳过
// 当 sum == 0 时，nums[L] == nums[L+1] 则会导致结果重复，应该跳过，L++
// 当 sum == 0 时，nums[R] == nums[R-1] 则会导致结果重复，应该跳过，R--
// 时间复杂度：O(n^2)，n为数组长度

func main() {
	arr := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(arr)
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}

func threeSum(nums []int) [][]int {
	if nums == nil || len(nums) < 3 {
		return nil
	}
	// 先排序
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 { //如果当前数字大于0，3个数相加一定大于零
			break
		}
		if i > 0 && nums[i] == nums[i-1] { //去重
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++ //去重
				}
				for left < right && nums[right] == nums[right-1] {
					right-- //去重
				}
				right--
				left++
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return res
}
