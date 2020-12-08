package main

// https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal
// 根据一棵树的中序遍历与后序遍历构造二叉树。

// 注意:
// 你可以假设树中没有重复的元素。

// 例如，给出

// 中序遍历 inorder = [9,3,15,20,7]
// 后序遍历 postorder = [9,15,7,20,3]
// 返回如下的二叉树：

//     3
//    / \
//   9  20
//     /  \
//    15   7

// Definition for a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	return build(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1)
}

func build(inorder []int, inStart, inEnd int, postorder []int, postStart, postEnd int) *TreeNode {
	// base case
	if inStart > inEnd {
		return nil
	}
	rootValue := postorder[postEnd]
	index := 0
	for i := inStart; i <= inEnd; i++ {
		if inorder[i] == rootValue {
			index = i
			break
		}
	}
	root := &TreeNode{Val: rootValue}
	leftSize := index - inStart
	root.Left = build(inorder, inStart, index-1, postorder, postStart, postStart+leftSize-1)
	root.Right = build(inorder, index+1, inEnd, postorder, postStart+leftSize, postEnd-1)
	return root
}

func main() {

}
