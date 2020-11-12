package tree

// BST 二叉排序树
type BST struct {
	Left  *BST
	Right *BST
	Value int
}

// Search 查找
func (b *BST) Search(v int) bool {
	if b == nil {
		return false
	}
	if b.Value == v {
		return true
	} else if b.Value > v {
		return b.Left.Search(v)
	}
	return b.Right.Search(v)
}

// Insert 插入元素
func (b *BST) Insert(v int) *BST {
	if b == nil {
		newNode := &BST{nil, nil, v}
		return newNode
	}
	if b.Value > v {
		b.Left = b.Left.Insert(v)
	} else {
		b.Right = b.Right.Insert(v)
	}
	return b
}

// Delete 删除元素
// 1、如果被删除节点只有一个子节点，就直接将A的子节点连至A的父节点上，并将A删除；
// 2、如果被删除节点有两个子节点，将节点右子树内的最小节点取代A。
func (b *BST) Delete(v int) *BST {
	if b == nil {
		return b
	}
	compare := v - b.Value
	if compare > 0 {
		b.Right = b.Right.Delete(v)
	} else if compare < 0 {
		b.Left = b.Left.Delete(v)
	} else { //找到节点，删除节点
		if b.Left != nil && b.Right != nil {
			b.Value = b.Right.getMin()
			b.Right = b.Right.Delete(b.Value)
		} else if b.Left != nil {
			b = b.Left
		} else {
			b = b.Right
		}
	}
	return b
}

// 按顺序获取树中元素
func (b *BST) getAll() []int {
	values := []int{}
	return addValues(values, b)
}

// 将一个节点加入到切片中
func addValues(values []int, b *BST) []int {
	if b != nil {
		values = addValues(values, b.Left)
		values = append(values, b.Value)
		values = addValues(values, b.Right)
	}
	return values
}

// 查找最大节点
func (b *BST) getMaxNode() *BST {
	if b == nil {
		return nil
	}
	for b.Right != nil {
		b = b.Right
	}
	return b
}

// 查找最小节点
func (b *BST) getMinNode() *BST {
	if b == nil {
		return nil
	}
	for b.Left != nil {
		b = b.Left
	}
	return b
}

// 获取最大值
func (b *BST) getMax() int {
	if b == nil {
		return -1
	}
	if b.Right == nil {
		return b.Value
	}
	return b.Right.getMax()
}

// 获取最小值
func (b *BST) getMin() int {
	if b == nil {
		return -1
	}
	if b.Left == nil {
		return b.Value
	}
	return b.Left.getMin()
}
