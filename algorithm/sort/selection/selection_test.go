package selection

import "testing"

func TestSelectionDesc(t *testing.T) {
	src := []int{2, 4, 1, 21, 123, 45, 7, 9, 12, 33}
	SelectionDesc(src)
	t.Log(src)
}

func TestSelectionAsc(t *testing.T) {
	src := []int{2, 4, 1, 21, 123, 45, 7, 9, 12, 33}
	SelectionAsc(src)
	t.Log(src)
}
