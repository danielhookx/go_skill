package _103

//103. Binary Tree Zigzag Level Order Traversal
//Given the root of a binary tree, return the zigzag level order traversal of its nodes' values. (i.e., from left to right, then right to left for the next level and alternate between).
//
//link: https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	order := false
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
		if order {
			revert(list)
		}
		order = !order
		result = append(result, list)
	}
	return result
}

func revert(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}
