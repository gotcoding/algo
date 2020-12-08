package main

// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal
// 根据一棵树的前序遍历与中序遍历构造二叉树。

// 注意:
// 你可以假设树中没有重复的元素。

// 例如，给出

// 前序遍历 preorder = [3,9,20,15,7]
// 中序遍历 inorder = [9,3,15,20,7]
// 返回如下的二叉树：

//     3
//    / \
//   9  20
//     /  \
//    15   7

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return build(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func build(preorder []int, preStart, preEnd int, inorder []int, inStart, inEnd int) *TreeNode {
	if preStart > preEnd {
		return nil
	}
	rootVal := preorder[preStart]
	index := 0
	for i := inStart; i <= inEnd; i++ {
		if inorder[i] == rootVal {
			index = i
			break
		}
	}
	leftSize := index - inStart
	root := &TreeNode{Val: rootVal}
	root.Left = build(preorder, preStart+1, preStart+leftSize, inorder, inStart, index-1)
	root.Right = build(preorder, preStart+leftSize+1, preEnd, inorder, index+1, inEnd)
	return root
}
