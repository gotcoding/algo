package main

import "math"

// 最大二叉树 https://leetcode-cn.com/problems/maximum-binary-tree
// 给定一个不含重复元素的整数数组。一个以此数组构建的最大二叉树定义如下：

// 二叉树的根是数组中的最大元素。
// 左子树是通过数组中最大值左边部分构造出的最大二叉树。
// 右子树是通过数组中最大值右边部分构造出的最大二叉树。
// 通过给定的数组构建最大二叉树，并且输出这个树的根节点。

// 示例 ：
// 输入：[3,2,1,6,0,5]
// 输出：返回下面这棵树的根节点：
//       6
//     /   \
//    3     5
//     \    /
//      2  0
//        \
//         1

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return build(nums, 0, len(nums)-1)
}

func build(nums []int, lo, hi int) *TreeNode {
	if lo > hi {
		return nil
	}
	// 找出数组中最大值对应的索引
	max := math.MinInt64
	index := -1
	for i := lo; i <= hi; i++ {
		if nums[i] > max {
			max = nums[i]
			index = i
		}
	}

	root := &TreeNode{Val: max, Left: nil, Right: nil}
	// 递归调用构建左右子树
	root.Left = build(nums, lo, index-1)
	root.Right = build(nums, index+1, hi)
	return root
}
