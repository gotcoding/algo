package main

import (
	"strconv"
)

// https://leetcode-cn.com/problems/find-duplicate-subtrees
// 给定一棵二叉树，返回所有重复的子树。对于同一类的重复子树，你只需要返回其中任意一棵的根结点即可。

// 两棵树重复是指它们具有相同的结构以及相同的结点值。

// 示例 1：
//         1
//        / \
//       2   3
//      /   / \
//     4   2   4
//        /
//       4
// 下面是两个重复的子树：

//       2
//      /
//     4
// 和
//     4
// 因此，你需要以列表的形式返回上述重复子树的根结点。

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	res, _ := doFindDuplicateSubtrees(root, make([]*TreeNode, 0), make(map[string]int))
	return res
}

func doFindDuplicateSubtrees(root *TreeNode, ans []*TreeNode, m map[string]int) ([]*TreeNode, string) {
	if root == nil {
		return ans, "*"
	}
	var l, r string
	ans, l = doFindDuplicateSubtrees(root.Left, ans, m)
	ans, r = doFindDuplicateSubtrees(root.Right, ans, m)
	s := l + "," + r + "," + strconv.Itoa(root.Val)
	m[s]++
	if m[s] == 2 {
		ans = append(ans, root)
	}
	return ans, s
}
