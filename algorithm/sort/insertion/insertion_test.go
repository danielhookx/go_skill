package insertion

import "testing"

func TestInsertionDesc(t *testing.T) {
	src := []int{2, 4, 1, 21, 123, 45, 7, 9, 12, 33}
	InsertionDesc(src)
	t.Log(src)
}

func TestInsertionAsc(t *testing.T) {
	src := []int{2, 4, 1, 21, 123, 45, 7, 9, 12, 33}
	InsertionAsc(src)
	t.Log(src)
}
