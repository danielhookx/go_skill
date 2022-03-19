package binary_tree

import "fmt"

type LinkedTreeNode struct {
	val   string
	left  *LinkedTreeNode
	right *LinkedTreeNode
}

func preOrder(n *LinkedTreeNode) {
	if n == nil {
		return
	}
	fmt.Println(n.val)
	preOrder(n.left)
	preOrder(n.right)
}

func inOrder(n *LinkedTreeNode) {
	if n == nil {
		return
	}
	inOrder(n.left)
	fmt.Println(n.val)
	inOrder(n.right)
}

func postOrder(n *LinkedTreeNode) {
	if n == nil {
		return
	}
	postOrder(n.left)
	postOrder(n.right)
	fmt.Println(n.val)
}
