package main

// https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node
// 给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：
// 填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
// 初始状态下，所有 next 指针都被设置为 NULL。

// 进阶：
// 你只能使用常量级额外空间。
// 使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。

// Definition for a Node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 递归法
// 时间复杂度O(n)
// 空间复杂度O(1)，不含递归栈占用
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	connectTwoNode(root.Left, root.Right)
	return root
}

func connectTwoNode(left, right *Node) {
	if left == nil || right == nil {
		return
	}
	// 前序遍历位置：将两个节点连接
	left.Next = right
	// 将左右节点内部连接
	connectTwoNode(left.Left, left.Right)
	connectTwoNode(right.Left, right.Right)
	// 实现跨左右节点连接
	connectTwoNode(left.Right, right.Left)
}

// 迭代法 DFS 按层遍历
func connect2(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		size := len(queue)
		// 依次将队列内的节点进行连接
		for i := 1; i < size; i++ {
			queue[i-1].Next = queue[i]
		}
		for i := 0; i < size; i++ {
			tmp := queue[0]
			queue = queue[1:]
			if tmp.Left != nil {
				queue = append(queue, tmp.Left)
			}
			if tmp.Right != nil {
				queue = append(queue, tmp.Right)
			}
		}
	}
	return root
}

func main() {

}
