package tree

import "fmt"

// Node 节点
type Node struct {
	left  *Node
	right *Node
	value interface{}
}

// CreateNode 创建节点
func CreateNode(v interface{}) *Node {
	return &Node{nil, nil, v}
}

// SetValue 赋值
func (n *Node) SetValue(v interface{}) {
	if n == nil {
		fmt.Println("Node is nil")
		return
	}
	n.value = v
}

// Print 打印
func (n *Node) Print() {
	fmt.Printf("%s ", n.value)
}

// preOrderTraversal 前序遍历
func (n *Node) preOrderTraversal() {
	if n == nil {
		return
	}
	n.Print()
	n.left.preOrderTraversal()
	n.right.preOrderTraversal()
}

// preOrderTraversalNew1
func (n *Node) preOrderTraversalNew1() []interface{} {
	if n == nil {
		return nil
	}
	if n.left == nil && n.right == nil {
		return []interface{}{n.value}
	}
	var res []interface{}
	res = append(res, n.value)
	res = append(res, n.left.preOrderTraversalNew1()...)
	res = append(res, n.right.preOrderTraversalNew1()...)
	return res
}

// preOrderTraversalNew
func (n *Node) preOrderTraversalNew() []interface{} {
	if n == nil {
		return nil
	}
	if n.left == nil && n.right == nil {
		return []interface{}{n.value}
	}
	var stack []*Node
	var res []interface{}
	stack = append(stack, n)
	for len(stack) > 0 {
		e := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, e.value)
		if e.right != nil {
			stack = append(stack, e.right)
		}
		if e.left != nil {
			stack = append(stack, e.left)
		}
	}
	return res
}

// midOrderTraversal 中序遍历
func (n *Node) midOrderTraversal() {
	if n == nil {
		return
	}
	n.left.midOrderTraversal()
	n.Print()
	n.right.midOrderTraversal()
}

// midOrderTraversalNew 中序遍历
func (n *Node) midOrderTraversalNew() []interface{} {
	if n == nil {
		return nil
	}
	if n.left == nil && n.right == nil {
		return []interface{}{n.value}
	}
	res := n.left.midOrderTraversalNew()
	res = append(res, n.value)
	res = append(res, n.right.midOrderTraversalNew()...)
	return res
}

func (n *Node) postOrderTraversal() {
	if n == nil {
		return
	}
	n.left.postOrderTraversal()
	n.right.postOrderTraversal()
	n.Print()
}

func (n *Node) postOrderTraversalNew() []interface{} {
	if n == nil {
		return nil
	}
	if n.left == nil && n.right == nil {
		return []interface{}{n.value}
	}
	res := n.left.postOrderTraversalNew()
	res = append(res, n.right.postOrderTraversalNew()...)
	res = append(res, n.value)
	return res
}

func (n *Node) breadthFirstTraversal() []interface{} {
	if n == nil {
		return nil
	}
	var res []interface{}

	nodes := []*Node{n}
	for len(nodes) > 0 {
		cur := nodes[0]
		nodes = nodes[1:]
		res = append(res, cur.value)
		if cur.left != nil {
			nodes = append(nodes, cur.left)
		}
		if cur.right != nil {
			nodes = append(nodes, cur.right)
		}
	}
	return res
}
