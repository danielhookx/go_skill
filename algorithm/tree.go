package algorithm

import (
	"errors"
	"strconv"
)

type TreeNodeI[T Any] interface {
	Val() T
	Left() TreeNodeI[T]
	Right() TreeNodeI[T]
}

func PreOrder[T Any](n TreeNodeI[T], rlt *[]T) {
	if IsNil[TreeNodeI[T]](n) {
		return
	}
	*rlt = append(*rlt, n.Val())
	PreOrder(n.Left(), rlt)
	PreOrder(n.Right(), rlt)
}

func InOrder[T Any](n TreeNodeI[T], rlt *[]T) {
	if IsNil[TreeNodeI[T]](n) {
		return
	}
	InOrder(n.Left(), rlt)
	*rlt = append(*rlt, n.Val())
	InOrder(n.Right(), rlt)
}

func PostOrder[T Any](n TreeNodeI[T], rlt *[]T) {
	if IsNil[TreeNodeI[T]](n) {
		return
	}
	PostOrder(n.Left(), rlt)
	PostOrder(n.Right(), rlt)
	*rlt = append(*rlt, n.Val())
}

func TreeBFS[T Any](root TreeNodeI[T]) [][]T {
	if IsNil[TreeNodeI[T]](root) {
		return nil
	}

	queue := NewQueue[TreeNodeI[T]]()
	queue.EnQueue(root)
	result := make([][]T, 0)

	for queue.Len() > 0 {
		l := queue.Len()
		level := make([]T, 0)
		for j := 0; j < l; j++ {
			node := queue.DeQueue()
			level = append(level, node.Val())

			if !IsNil[TreeNodeI[T]](node.Left()) {
				queue.EnQueue(node.Left())
			}

			if !IsNil[TreeNodeI[T]](node.Right()) {
				queue.EnQueue(node.Right())
			}
		}
		result = append(result, level)
	}
	return result
}

type TreeNode[T Any] struct {
	val   T
	left  *TreeNode[T]
	right *TreeNode[T]
}

func (t *TreeNode[T]) Val() T {
	return t.val
}

func (t *TreeNode[T]) Left() TreeNodeI[T] {
	return t.left
}

func (t *TreeNode[T]) Right() TreeNodeI[T] {
	return t.right
}

func ParseFunc[T Any]() func(val string) (T, error) {
	var t T
	switch Any(t).(type) {
	case int:
		return func(val string) (T, error) {
			v, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return Any(0).(T), err
			}
			return Any(int(v)).(T), nil
		}
	case string:
		return func(val string) (T, error) {
			if val == "" || val == "nil" {
				return Any(val).(T), errors.New("nil")
			}
			return Any(val).(T), nil
		}
	default:
		return func(val string) (T, error) {
			panic("unsupported type")
		}
	}
}

func NewTree[T Any](level ...[]string) *TreeNode[T] {
	sum := make([]string, 0)
	for _, l := range level {
		sum = append(sum, l...)
	}
	if len(level) < 1 || len(level[0]) < 1 {
		return nil
	}
	parse := ParseFunc[T]()
	v, err := parse(level[0][0])
	if err != nil {
		return nil
	}
	root := &TreeNode[T]{
		val:   v,
		left:  nil,
		right: nil,
	}

	queue := NewQueue[*TreeNode[T]]()
	queue.EnQueue(root)
	indexQueue := NewQueue[int]()
	indexQueue.EnQueue(0)

	for queue.Len() > 0 {
		l := queue.Len()
		for i := 0; i < l; i++ {
			node := queue.DeQueue()
			index := indexQueue.DeQueue()

			leftChildIndex := index*2 + 1
			rightChildIndex := index*2 + 2
			if leftChildIndex < len(sum) {
				lVal, err := parse(sum[leftChildIndex])
				if err == nil {
					node.left = &TreeNode[T]{
						val:   lVal,
						left:  nil,
						right: nil,
					}
					queue.EnQueue(node.left)
					indexQueue.EnQueue(leftChildIndex)
				}
			}

			if rightChildIndex < len(sum) {
				rVal, err := parse(sum[rightChildIndex])
				if err == nil {
					node.right = &TreeNode[T]{
						val:   rVal,
						left:  nil,
						right: nil,
					}
					queue.EnQueue(node.right)
					indexQueue.EnQueue(rightChildIndex)
				}
			}
		}
	}
	return root
}
