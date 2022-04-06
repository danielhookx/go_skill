package merge

import "testing"

func TestMergeSort(t *testing.T) {
	src := []int{2, 4, 1, 21, 123, 45, 7, 9, 12, 33}
	MergeSort(src)
	t.Log(src)
}
