package binary_tree

import (
	"testing"
)

var head *LinkedTreeNode

func Test_putSort(t *testing.T) {
	src := []int{3, 1, 2, 5, 9, 6, 4, 8}
	var h = &LinkedTreeNode{
		val:   0,
		left:  nil,
		right: nil,
	}
	for _, data := range src {
		putSort(h, data)
	}
	inOrder(h)
}

func Test_putSort2(t *testing.T) {
	src := []int{3, 1, 2, 5, 9, 6, 4, 8}
	var h = &LinkedTreeNode{
		val:   0,
		left:  nil,
		right: nil,
	}
	for _, data := range src {
		putSort2(h, data)
	}
	inOrder(h)
}
func Test_search(t *testing.T) {
	src := []int{3, 1, 2, 5, 9, 6, 4, 8}
	var h = &LinkedTreeNode{
		val:   0,
		left:  nil,
		right: nil,
	}
	for _, data := range src {
		putSort2(h, data)
	}
	inOrder(h)
}
