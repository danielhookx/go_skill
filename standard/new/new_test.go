package new

import "testing"

func Test_New(t *testing.T) {
	var sPrt *[]int
	sPrt = new([]int)
	if *sPrt == nil {
		t.Log("sPtr is nil")
	} else {
		t.Log("sPtr is empty")
	}

	var s []int
	s = make([]int, 0)
}
