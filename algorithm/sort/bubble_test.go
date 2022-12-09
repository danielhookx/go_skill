package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortAsc(t *testing.T) {
	src := []int{2, 4, 1, 21, 123, 45, 7, 9, 12, 33}
	BubbleSort(src)
	assert.EqualValues(t, []int{1, 2, 4, 7, 9, 12, 21, 33, 45, 123}, src)
}

func BenchmarkBubbleSort(b *testing.B) {
	src := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	BubbleSort(src)
}

func BenchmarkBubbleSortFlag(b *testing.B) {
	src := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	BubbleSortFlag(src)
}
