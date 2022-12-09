package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	src := []int{2, 4, 1, 21, 123, 45, 7, 9, 12, 33}
	MergeSort(src)
	assert.EqualValues(t, []int{1, 2, 4, 7, 9, 12, 21, 33, 45, 123}, src)
}
