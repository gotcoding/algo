package tree

import (
	"errors"
	"fmt"
)

// AVLNode 二叉平衡树
type AVLNode struct {
	Left, Right *AVLNode    // 左右子节点
	height      int         // 节点高度
	Data        interface{} // 数据
}

// 比较指针
var compare Comparator

// Comparator 指针类型比较，用于比较节点内数值大小
type Comparator func(a, b interface{}) int

// NewAVLNode 新建AVL节点
func NewAVLNode(v interface{}) *AVLNode {
	return &AVLNode{nil, nil, 1, v}
}

// NewAVLTree 新建平衡二叉树
func NewAVLTree(data interface{}, myfunc Comparator) (*AVLNode, error) {
	if data == nil || myfunc == nil {
		return nil, errors.New("不能为空")
	}
	compare = myfunc
	return NewAVLNode(data), nil
}

// GetData 获取节点数据
func (n *AVLNode) GetData() interface{} {
	return n.Data
}

// SetData 给节点赋值
func (n *AVLNode) SetData(data interface{}) {
	if n == nil {
		return
	}
	n.Data = data
}

// GetLeft 获取左节点
func (n *AVLNode) GetLeft() *AVLNode {
	if n == nil {
		return nil
	}
	return n.Left
}

// GetRight 获取右节点
func (n *AVLNode) GetRight() *AVLNode {
	if n == nil {
		return nil
	}
	return n.Right
}

// GetHeight 获取高度
func (n *AVLNode) GetHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

// Max 比较两个子树高度大小
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Find 查找节点
func (n *AVLNode) Find(data interface{}) *AVLNode {
	var findedNode *AVLNode
	switch compare(data, n.Data) {
	case -1:
		findedNode = n.Left.Find(data)
	case 1:
		findedNode = n.Right.Find(data)
	case 0:
		return n
	}
	return findedNode
}

// FindMax 查找最大值
func (n *AVLNode) FindMax() *AVLNode {
	var findedNode *AVLNode
	if n == nil {
		return nil
	}
	if n.Right != nil {
		findedNode = n.Right.FindMax()
	} else {
		findedNode = n
	}
	return findedNode
}

// FindMin 查找最小值
func (n *AVLNode) FindMin() *AVLNode {
	var findedNode *AVLNode
	if n == nil {
		return nil
	}
	if n.Left != nil {
		findedNode = n.Left.FindMin()
	} else {
		findedNode = n
	}
	return findedNode
}

// RightRotate 顺时针旋转，右旋
func (n *AVLNode) RightRotate() *AVLNode {
	headNode := n.Left
	n.Left = headNode.Right
	headNode.Right = n

	// 更新旋转后的节点高度
	n.height = Max(n.Left.Right.height, n.Right.height) + 1
	headNode.height = Max(headNode.Left.height, headNode.Right.height) + 1
	return headNode
}

// LeftRotate 逆时针旋转，左旋
func (n *AVLNode) LeftRotate() *AVLNode {
	headNode := n.Right
	n.Right = headNode.Left
	headNode.Left = n

	// 更新旋转后的节点高度
	n.height = Max(n.Left.Right.height, n.Right.height) + 1
	headNode.height = Max(headNode.Left.height, headNode.Right.height) + 1
	return headNode
}

// LeftThenRightRotate 先逆时针旋转再顺时针旋转，先左旋在再右旋
func (n *AVLNode) LeftThenRightRotate() *AVLNode {
	// 先把左子节点左旋转
	n.Left = n.Left.LeftRotate()
	// 再把自己右旋转
	return n.RightRotate()
}

// RightThenLeftRotate 先顺时针旋转再逆时针旋转，先右旋再左旋
func (n *AVLNode) RightThenLeftRotate() *AVLNode {
	// 先把右子节点右旋
	n.Right = n.Right.RightRotate()
	// 再把自己左旋转
	return n.LeftRotate()
}

// 调整AVL树
func (n *AVLNode) adjust() *AVLNode {
	// 如果右子树的高度比左子树高大于2
	if n.Right.GetHeight()-n.Left.GetHeight() == 2 {
		// 如果n.Right 的右子树的高度比n.Left 的左子树高度大,直接对n进行左旋
		// 否则先对n.Right进行右旋然后再对n进行左旋
		if n.Right.Right.GetHeight() > n.Right.Left.GetHeight() {
			return n.LeftRotate()
		}
		n.Right = n.Right.RightRotate()
		return n.LeftRotate()
	} else if n.Right.GetHeight()-n.Left.GetHeight() == -2 {
		// 如果n.Left 的左子树的高度比n.Right 的右子树高度大,直接对n进行右旋
		// 否则先对n.Left进行左旋然后再对n进行右旋
		if n.Left.Left.GetHeight() > n.Left.Right.GetHeight() {
			return n.RightRotate()
		}
		n.Left = n.Left.LeftRotate()
		return n.RightRotate()
	}
	return n
}

// Insert 插入节点
// 递归插入，需要调整树的平衡和高度
func (n *AVLNode) Insert(data interface{}) *AVLNode {
	if n == nil {
		return NewAVLNode(data)
	}
	switch compare(data, n.GetData()) {
	case -1: //左子树插入
		n.Left = n.Left.Insert(data)
		n = n.adjust()
	case 1: // 右子树插入
		n.Right = n.Right.Insert(data)
		n = n.adjust()
	case 0:
		fmt.Print("数据已存在")
	}
	n.height = Max(n.Left.GetHeight(), n.Right.GetHeight()) + 1
	return n
}

// Delete 删除节点
func (n *AVLNode) Delete(data interface{}) *AVLNode {
	if n == nil {
		return nil
	}
	switch compare(data, n.Data) {
	case -1:
		n.Right = n.Right.Delete(data)
	case 1:
		n.Left = n.Left.Delete(data)
	case 0: //找到数据，删除节点
		if n.Left != nil && n.Right != nil { //左右节点均有值
			n.Data = n.Right.FindMin().Data
			n.Right = n.Right.Delete(data)
		} else if n.Left != nil { //右节点有值
			n = n.Left
		} else {
			n = n.Right
		}
	}
	// 自动调整高度
	if n != nil {
		n.height = Max(n.Left.GetHeight(), n.Right.GetHeight()) + 1
		n = n.adjust()
	}
	return n
}
