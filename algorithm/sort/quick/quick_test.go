package quick

import "testing"

func TestQuickSort(t *testing.T) {
	src := []int{2, 4, 1, 21, 123, 45, 7, 9, 12, 33}
	QuickSort(src, 0, len(src))
	t.Log(src)
}
