package tree

import (
	"fmt"
	"testing"
)

func TestPreOrder(t *testing.T) {
	root := CreateNode("A")
	root.left = CreateNode("B")
	root.left.left = CreateNode("D")
	root.left.right = CreateNode("E")
	root.left.left.left = CreateNode("H")
	root.left.left.right = CreateNode("I")
	root.left.right.left = CreateNode("J")
	root.right = CreateNode("C")
	root.right.right = CreateNode("G")
	root.right.left = CreateNode("F")

	root.preOrderTraversal()
	res := root.preOrderTraversalNew()
	res1 := root.preOrderTraversalNew1()
	fmt.Println(res)
	fmt.Println(res1)
	root.midOrderTraversal()
	midRes := root.midOrderTraversalNew()
	fmt.Println(midRes)
	root.postOrderTraversal()
	fmt.Println(root.postOrderTraversalNew())

	fmt.Println(root.breadthFirstTraversal())
	fmt.Println(root.Layers())
	fmt.Printf("查找：%s\n", root.FindNode(root, "H").value)
}
