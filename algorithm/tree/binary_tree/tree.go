package binary_tree

import (
	"fmt"
	"github.com/oofpgDLD/go_skill/algorithm/queue"
)

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

func BFSByCycle(root *LinkedTreeNode) []string {
	var temp_node *LinkedTreeNode
	result := make([]string, 0)
	que := queue.NewQueue()
	temp_node = root
	for temp_node != nil {
		result = append(result, temp_node.val)
		que.Enqueue(temp_node.left)
		que.Enqueue(temp_node.right)
		temp_node = que.Dequeue().(*LinkedTreeNode)
	}
	return result
}

func DFSByRecursively(root *LinkedTreeNode) []string {
	result := make([]string, 0)
	dfs(root, &result)
	return result
}

func dfs(root *LinkedTreeNode, result *[]string) {
	if root == nil {
		return
	}
	*result = append(*result, root.val)
	dfs(root.left, result)
	dfs(root.right, result)
}
