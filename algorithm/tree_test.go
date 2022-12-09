package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPreOrder(t *testing.T) {
	l1 := []string{"6"}
	l2 := []string{"1", "3"}
	l3 := []string{"9", "4", "2", "7"}
	l4 := []string{"5", "nil", "nil", "nil", "8", "nil", "nil", "nil"}
	root := NewTree(l1, l2, l3, l4)
	rlt := make([]int, 0)
	PreOrder(root, &rlt)
	assert.EqualValues(t, []int{6, 1, 9, 5, 4, 3, 2, 8, 7}, rlt)
}

func TestInOrder(t *testing.T) {
	l1 := []string{"6"}
	l2 := []string{"1", "3"}
	l3 := []string{"9", "4", "2", "7"}
	l4 := []string{"5", "nil", "nil", "nil", "8", "nil", "nil", "nil"}
	root := NewTree(l1, l2, l3, l4)
	rlt := make([]int, 0)
	InOrder(root, &rlt)
	assert.EqualValues(t, []int{5, 9, 1, 4, 6, 8, 2, 3, 7}, rlt)
}

func TestPostOrder(t *testing.T) {
	l1 := []string{"6"}
	l2 := []string{"1", "3"}
	l3 := []string{"9", "4", "2", "7"}
	l4 := []string{"5", "nil", "nil", "nil", "8", "nil", "nil", "nil"}
	root := NewTree(l1, l2, l3, l4)
	rlt := make([]int, 0)
	PostOrder(root, &rlt)
	assert.EqualValues(t, []int{5, 9, 4, 1, 8, 2, 7, 3, 6}, rlt)
}

func TestBFS(t *testing.T) {
	l1 := []string{"6"}
	l2 := []string{"1", "3"}
	l3 := []string{"9", "4", "2", "7"}
	l4 := []string{"5", "nil", "nil", "nil", "8", "nil", "nil", "nil"}
	root := NewTree(l1, l2, l3, l4)
	rlt := BFS(root)
	assert.EqualValues(t, [][]int{
		{6},
		{1, 3},
		{9, 4, 2, 7},
		{5, 8},
	}, rlt)
}

func TestSearchTree(t *testing.T) {
	st := &SearchTree{}
	src := []int{5, 9, 4, 1, 8, 2, 7, 3, 6}
	for _, v := range src {
		st.Put(v)
	}
	inOrderRlt := make([]int, 0)
	InOrder(st.root, &inOrderRlt)
	assert.EqualValues(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, inOrderRlt)
}
