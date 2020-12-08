package main

// https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list
// 二叉树展开为链表
// 给定一个二叉树，原地将它展开为一个单链表。

// 例如，给定二叉树

//     1
//    / \
//   2   5
//  / \   \
// 3   4   6
// 将其展开为：

// 1
//  \
//   2
//    \
//     3
//      \
//       4
//        \
//         5
//          \
//           6

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)
	// 左右节点交换
	root.Left, root.Right = root.Right, root.Left
	// 将左子节点移到右子节点的末端
	left := root.Left
	root.Left = nil
	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = left
}

func flatten2(root *TreeNode) {
	cur := root
	for cur != nil {
		if cur.Left != nil {
			next := cur.Left
			pre := next
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right = cur.Right
			cur.Left, cur.Right = nil, next
		}
		cur = cur.Right
	}
}

func main() {

}
