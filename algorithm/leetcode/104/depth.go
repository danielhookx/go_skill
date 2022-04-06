package _104

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
