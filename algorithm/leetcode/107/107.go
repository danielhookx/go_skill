package _107

//107. Binary Tree Level Order Traversal II
//Given the root of a binary tree, return the bottom-up level order traversal of its nodes' values. (i.e., from left to right, level by level from leaf to root).
//
//link:https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	results := levelOrder(root)
	reverse(results)
	return results
}

func reverse(nums [][]int) {
	if len(nums) < 1 {
		return
	}
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
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
