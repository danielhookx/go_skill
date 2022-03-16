package array

import "testing"

func Test_InitArray(t *testing.T) {
	a1 := [3]int{1, 2, 3}
	a2 := [...]int{1, 2, 3}
	var a3 [3]int

	t.Log(a1, a2, a3)
}
