package bubble

import "testing"

func TestBubbleSortDesc(t *testing.T) {
	src := []int{2, 4, 1, 21, 123, 45, 7, 9, 12, 33}
	BubbleSortDesc(src)
	t.Log(src)
}

func TestBubbleSortAsc(t *testing.T) {
	src := []int{2, 4, 1, 21, 123, 45, 7, 9, 12, 33}
	BubbleSortAsc(src)
	t.Log(src)
}
