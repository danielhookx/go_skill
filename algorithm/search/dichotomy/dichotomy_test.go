package dichotomy

import "testing"

func TestCycle(t *testing.T) {
	src := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 50, 100}
	t.Log(Cycle(src, 3))
	t.Log(Cycle(src, 50))
	t.Log(Cycle(src, 21))
}

func Test_recursively(t *testing.T) {
	src := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 50, 100}
	t.Log(recursively(src, 0, len(src)-1, 3))
	t.Log(recursively(src, 0, len(src)-1, 50))
	t.Log(recursively(src, 0, len(src)-1, 21))
}
