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
	_, lfmax := maxPath(root)
	return lfmax
}

func maxPath(root *TreeNode) (int, int) {
	if root == nil {
		return 0, -(1 << 31)
	}
	//value 1 is relatively max, value 2 is final max
	lramx, lfmax := maxPath(root.Left)
	rramx, rfmax := maxPath(root.Right)

	singlePath := max(max(lramx, rramx)+root.Val, 0)
	maxpath := max(max(lfmax, rfmax), root.Val+lramx+rramx)
	return singlePath, maxpath
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
