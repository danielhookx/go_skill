package binary_tree

import (
	"os"
	"testing"
)

var head *LinkedTreeNode

func TestMain(t *testing.M) {
	head = &LinkedTreeNode{
		val:   "A",
		left:  &LinkedTreeNode{
			val:   "B",
			left:  &LinkedTreeNode{
				val:   "D",
				left:  nil,
				right: nil,
			},
			right: &LinkedTreeNode{
				val:   "E",
				left:  nil,
				right: nil,
			},
		},
		right: &LinkedTreeNode{
			val:   "C",
			left:  &LinkedTreeNode{
				val:   "F",
				left:  nil,
				right: nil,
			},
			right: &LinkedTreeNode{
				val:   "G",
				left:  nil,
				right: nil,
			},
		},
	}
	os.Exit(t.Run())
}

func Test_inOrder(t *testing.T) {
	inOrder(head)
}

func Test_postOrder(t *testing.T) {
	postOrder(head)
}

func Test_preOrder(t *testing.T) {
	preOrder(head)
}

func Test_putSort(t *testing.T) {
	src := []string{"3","1","2", "5", "9", "6", "4", "8"}
	var h = &LinkedTreeNode{
		val:   "0",
		left:  nil,
		right: nil,
	}
	for _, data := range src {
		putSort2(h, data)
	}
	inOrder(h)
}

func Test_search(t *testing.T) {
	src := []string{"3","1","2", "5", "9", "6", "4", "8"}
	var h = &LinkedTreeNode{
		val:   "0",
		left:  nil,
		right: nil,
	}
	for _, data := range src {
		putSort2(h, data)
	}
	inOrder(h)
}