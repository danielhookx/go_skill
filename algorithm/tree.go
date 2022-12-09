package algorithm

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTree(level ...[]string) *TreeNode {
	sum := make([]string, 0)
	for _, l := range level {
		sum = append(sum, l...)
	}
	if len(level) < 1 || len(level[0]) < 1 {
		return nil
	}
	v, err := strconv.ParseInt(level[0][0], 10, 64)
	if err != nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	indexQueue := make([]int, 0)
	root := &TreeNode{
		Val:   int(v),
		Left:  nil,
		Right: nil,
	}
	queue = append(queue, root)
	indexQueue = append(indexQueue, 0)

	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]

			index := indexQueue[0]
			indexQueue = indexQueue[1:]

			leftChildIndex := index*2 + 1
			rightClildIndex := index*2 + 2
			if leftChildIndex < len(sum) {
				lVal, err := strconv.ParseInt(sum[leftChildIndex], 10, 64)
				if err == nil {
					node.Left = &TreeNode{
						Val:   int(lVal),
						Left:  nil,
						Right: nil,
					}
					queue = append(queue, node.Left)
					indexQueue = append(indexQueue, leftChildIndex)
				}
			}

			if rightClildIndex < len(sum) {
				rVal, err := strconv.ParseInt(sum[rightClildIndex], 10, 64)
				if err == nil {
					node.Right = &TreeNode{
						Val:   int(rVal),
						Left:  nil,
						Right: nil,
					}
					queue = append(queue, node.Right)
					indexQueue = append(indexQueue, rightClildIndex)
				}
			}
		}
	}
	return root
}

func PreOrder(n *TreeNode, rlt *[]int) {
	if n == nil {
		return
	}
	*rlt = append(*rlt, n.Val)
	PreOrder(n.Left, rlt)
	PreOrder(n.Right, rlt)
}

func InOrder(n *TreeNode, rlt *[]int) {
	if n == nil {
		return
	}
	InOrder(n.Left, rlt)
	*rlt = append(*rlt, n.Val)
	InOrder(n.Right, rlt)
}

func PostOrder(n *TreeNode, rlt *[]int) {
	if n == nil {
		return
	}
	PostOrder(n.Left, rlt)
	PostOrder(n.Right, rlt)
	*rlt = append(*rlt, n.Val)
}

func BFS(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	result := make([][]int, 0)

	for len(queue) > 0 {
		l := len(queue)
		level := make([]int, 0)
		for j := 0; j < l; j++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
	}
	return result
}

type SearchTree struct {
	root *TreeNode
}

func (r *SearchTree) Put(data int) {
	if r.root == nil {
		r.root = &TreeNode{Val: data}
		return
	}
	r.sortPut(r.root, data)
}

func (r *SearchTree) sortPut(n *TreeNode, data int) {
	if n.Val == data {
		return
	}
	if n.Val > data {
		if n.Left == nil {
			n.Left = &TreeNode{Val: data}
			return
		}
		r.sortPut(n.Left, data)
	}
	if n.Val < data {
		if n.Right == nil {
			n.Right = &TreeNode{Val: data}
			return
		}
		r.sortPut(n.Right, data)
	}
}
