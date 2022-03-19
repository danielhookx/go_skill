package binary_tree

import "fmt"

type LinkedTreeNode struct {
	val string
	left *LinkedTreeNode
	right *LinkedTreeNode
}

func putSort(n *LinkedTreeNode, data string) {
	if n == nil {
		n = &LinkedTreeNode{
			val:   data,
			left:  nil,
			right: nil,
		}
		return
	}
	if n.val == data {
		return
	}
	if n.val < data {
		putSort(n.left, data)
	}
	if n.val > data {
		putSort(n.right, data)
	}
}

func putSort2(tree *LinkedTreeNode, data string) {
	if tree == nil {
		tree = &LinkedTreeNode{
			val:   data,
			left:  nil,
			right: nil,
		}
		return
	}

	var p = tree
	for p != nil{
		if p.val > data {
			if p.right == nil {
				p.right = &LinkedTreeNode{
					val:   data,
					left:  nil,
					right: nil,
				}
				return
			}
			p = p.right
		}else if p.val < data {
			if p.left == nil {
				p.left = &LinkedTreeNode{
					val:   data,
					left:  nil,
					right: nil,
				}
				return
			}
			p = p.left
		}
	}
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