package main

import (
	"fmt"
	"sort"
)

// 链接：https://leetcode-cn.com/problems/majority-element
// 给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
// 示例:
// 输入: [2,2,1,1,1,2,2]
// 输出: 2

func main() {
	arr := []int{2, 2, 1, 1, 2, 2, 1}
	fmt.Println(majorityElement1(arr))
	fmt.Println(majorityElement2(arr))
	fmt.Println(majorityElement3(arr))
}

// 方法一：哈希表
// 新建一个map来统计每个数字出现的次数
// 从map中获取多数元素
func majorityElement1(nums []int) int {
	dic := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		if v, ok := dic[nums[i]]; ok {
			dic[nums[i]] = v + 1
		} else {
			dic[nums[i]] = 1
		}
	}
	for k, v := range dic {
		if v > len(nums)/2 {
			return k
		}
	}
	return -1
}

// 方法二：排序
// 如果将数组 nums 中的所有元素按照单调递增或单调递减的顺序排序，那么下标为 n/2 的元素（下标从 0 开始）一定是众数。
// 时间复杂度O(nlogn)，空间复杂度O(1)
func majorityElement2(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// 方法三：Boyer-Moore 投票算法
// 如果我们把众数记为 +1+1，把其他数记为 -1−1，将它们全部加起来，显然和大于 0，从结果本身我们可以看出众数比其他数多。
// 时间复杂度O(n)，空间复杂度O(1)
func majorityElement3(nums []int) int {
	count, candidate := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == candidate {
			count++
		} else {
			if count == 0 {
				candidate = nums[i]
				continue
			}
			count--
		}
	}
	return candidate
}
