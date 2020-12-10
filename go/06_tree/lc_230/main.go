package main

// https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst
// 给定一个二叉搜索树，编写一个函数 kthSmallest 来查找其中第 k 个最小的元素。

// 说明：
// 你可以假设 k 总是有效的，1 ≤ k ≤ 二叉搜索树元素个数。

// 示例 1:

// 输入: root = [3,1,4,null,2], k = 1
//    3
//   / \
//  1   4
//   \
//    2
// 输出: 1

// Definition for a binary tree node.÷
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	arr := make([]int, 0)
	inOrder(root, &arr)
	return arr[k-1]
}

func inOrder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, arr)
	*arr = append(*arr, root.Val)
	inOrder(root.Right, arr)
}

func inorderTravel(root *TreeNode) (result []int) {
	if nil == root {
		return []int{}
	}

	left := inorderTravel(root.Left)
	right := inorderTravel(root.Right)
	result = append(result, left...)
	result = append(result, root.Val)
	result = append(result, right...)

	return result
}
func kthSmallest3(root *TreeNode, k int) int {
	arr := inorderTravel(root)
	// fmt.Printf("arr:%v\n", arr)
	return arr[k-1]
}

var res, rank int

// 迭代法
func kthSmallest2(root *TreeNode, k int) int {
	res = 0
	rank = 0
	traverse(root, k)
	return res
}

func traverse(root *TreeNode, k int) {
	if root == nil {
		return
	}
	traverse(root.Left, k)
	// 中序遍历：处理数据
	rank++
	if rank == k {
		res = root.Val
		return
	}
	traverse(root.Right, k)
}
