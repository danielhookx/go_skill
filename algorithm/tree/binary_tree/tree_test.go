package binary_tree

import (
	"os"
	"testing"
)

var head *LinkedTreeNode

func TestMain(t *testing.M) {
	head = &LinkedTreeNode{
		val: "A",
		left: &LinkedTreeNode{
			val: "B",
			left: &LinkedTreeNode{
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
			val: "C",
			left: &LinkedTreeNode{
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
