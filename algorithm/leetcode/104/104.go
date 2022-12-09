package _104

//Maximum Depth of Binary Tree
//Given the root of a binary tree, return its maximum depth.
//
//A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
//
//link:https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ldeep := maxDepth(root.Left)
	rdeep := maxDepth(root.Right)
	if ldeep > rdeep {
		return ldeep + 1
	} else {
		return rdeep + 1
	}
}
