package _102

//102. Binary Tree Level Order Traversal
//Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).
//
//link:https://leetcode-cn.com/problems/binary-tree-level-order-traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	results := make([][]int, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		l := len(queue)
		list := make([]int, l)
		for i := 0; i < l; i++ {
			//dequeue
			node := queue[0]
			queue = queue[1:]
			list[i] = node.Val
			//enqueue
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		results = append(results, list)
	}
	return results
}
