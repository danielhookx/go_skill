package _124

// Binary Tree Maximum Path Sum
//A path in a binary tree is a sequence of nodes where each pair of adjacent nodes in the sequence has an edge connecting them. A node can only appear in the sequence at most once. Note that the path does not need to pass through the root.
//
//The path sum of a path is the sum of the node's values in the path.
//
//Given the root of a binary tree, return the maximum path sum of any non-empty path.
//
//Link: https://leetcode-cn.com/problems/binary-tree-maximum-path-sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxsum := root.Val
	justLeft := false
	if root.Left != nil {
		lmps := maxPathSum(root.Left)
		if m := max(maxsum, maxsum+lmps); m < lmps {
			maxsum = lmps
			justLeft = true
		} else {
			maxsum = m
		}
	}

	if root.Right != nil {
		rmps := maxPathSum(root.Right)
		if justLeft {
			return max(maxsum, rmps)
		}
		maxsum = max(maxsum, max(maxsum+rmps, rmps))
	}
	return maxsum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
