package slice

import "testing"

func Test_InitSlice(t *testing.T) {
	a1 := [5]int{1, 2, 3, 4, 5}
	s1 := []int{1, 2, 3}
	s2 := a1[:]
	s3 := make([]int, 4, 10)

	t.Log(s1, s2, s3)
	t.Log(cap(a1), len(a1), cap(s2), len(s2))
}

func Test_InitSliceAppend(t *testing.T) {
	a1 := [5]int{1, 2, 3, 4, 5}
	s1 := a1[:2]
	t.Log(a1, s1)

	s2 := append(s1, 6, 7)
	t.Log(a1, s1, s2)
}

func Test_InitSliceAppend2(t *testing.T) {
	var arr []int64
	arr = append(arr, 1, 2, 3, 4, 5)
	t.Log(arr)
	t.Logf("len:%d, cap:%d", len(arr), cap(arr))
}

func Test_InitSliceAppend3(t *testing.T) {
	var arr []int32
	arr = append(arr, 1, 2, 3, 4, 5)
	t.Log(arr)
	t.Logf("len:%d, cap:%d", len(arr), cap(arr))
	t.Logf("%x,%x", ^uintptr(0), ^uintptr(0)>>63)
}
