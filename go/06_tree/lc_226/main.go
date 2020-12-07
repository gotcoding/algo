package main

// https://leetcode-cn.com/problems/invert-binary-tree
// 翻转一棵二叉树。

// 示例：
// 输入：

//      4
//    /   \
//   2     7
//  / \   / \
// 1   3 6   9
// 输出：

//      4
//    /   \
//   7     2
//  / \   / \
// 9   6 3   1

// Definition for a binary tree node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归方法
// 时间复杂度O(n)，因为需要遍历树
// 空间复杂度O(1)，递归的过程中只调整了指针
func invertTree(root *TreeNode) *TreeNode {
	// base case
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

// 递归方式
// 时间复杂度O(n)，因为需要遍历树
// 空间复杂度O(n)，递归过程中占用新的空间
func invertTree2(root *TreeNode) *TreeNode {
	// base case
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	rigth := invertTree(root.Right)
	root.Left = rigth
	root.Right = left
	return root
}

// 迭代方式 BFS 广度优先，用层序遍历的方式去遍历二叉树。
func invertTree3(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return root
}

func main() {

}
