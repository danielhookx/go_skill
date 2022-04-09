package _110

//Given a binary tree, determine if it is height-balanced.
//
//For this problem, a height-balanced binary tree is defined as:
//
//a binary tree in which the left and right subtrees of every node differ in height by no more than 1.
//
//link：https://leetcode-cn.com/problems/balanced-binary-tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if depth(root) < 0 {
		return false
	}
	return true
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ldeep := depth(root.Left)
	rdeep := depth(root.Right)
	if ldeep < 0 || rdeep < 0 {
		return -1
	}
	if ldeep > rdeep {
		if ldeep-rdeep > 1 {
			return -1
		}
		return ldeep + 1
	} else {
		if rdeep-ldeep > 1 {
			return -1
		}
		return rdeep + 1
	}
}
