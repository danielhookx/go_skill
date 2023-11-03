package slice

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestSliceInitialization(t *testing.T) {
	// Using slice literal:
	s1 := []int{1, 2, 3, 4, 5}

	//Using an Array:
	a1 := [5]int{1, 2, 3, 4, 5}
	s2 := a1[:]
	assert.Equal(t, s1, s2)
	// a1 pointed by s2
	assert.Equal(t, uintptr(unsafe.Pointer(&s2[0])), uintptr(unsafe.Pointer(&a1)))
	assert.NotEqual(t, uintptr(unsafe.Pointer(&s2)), uintptr(unsafe.Pointer(&a1)))

	//Using already Existing Slice:
	s3 := s2[:]
	assert.Equal(t, s2, s3)
	// a1 pointed by s3
	assert.Equal(t, uintptr(unsafe.Pointer(&s3[0])), uintptr(unsafe.Pointer(&a1)))
	assert.NotEqual(t, uintptr(unsafe.Pointer(&s2)), uintptr(unsafe.Pointer(&s3)))

	//Using make() function:
	s4 := make([]int, 4, 7)
	s5 := make([]int, 7)
	assert.Equal(t, 4, len(s4))
	assert.Equal(t, 7, cap(s4))
	assert.Equal(t, 7, len(s5))
	assert.Equal(t, 7, cap(s5))
}

func Test_SliceAppend(t *testing.T) {
	// insert into not init slice it`s ok
	var nilS []int
	nilS = append(nilS, 1)
	assert.Equal(t, 1, nilS[0])

	s := make([]int, 3)
	s1 := append(s, []int{1, 2, 3, 4}...)
	//s1:len=7, cap=8
	//new_cap = cap(64 / 8)
	// 7 * 8 byte = 56B
	//span level 6 = 64B
	assert.Equal(t, 7, len(s1))
	assert.Equal(t, 8, cap(s1))

	s = make([]int, 256)
	s1 = append(s, 1)
	//s1:len=257, cap=512
	//new_cap = double old_cap
	assert.Equal(t, 256+1, len(s1))
	assert.Equal(t, 256*2, cap(s1))

	s = make([]int, 1024)
	s1 = append(s, 1)
	//s1:len=1025, cap=1280
	//new_cap = old_cap * 1.25
	assert.Equal(t, 1024+1, len(s1))
	assert.Equal(t, int(1024*1.25), cap(s1))
}

func TestFuncParamsValueCopy(t *testing.T) {
	f := func(s []int) {
		s = append(s, []int{1, 2, 3, 4, 5}...)
	}
	f2 := func(s *[]int) {
		*s = append(*s, []int{1, 2, 3, 4, 5}...)
	}
	s := make([]int, 0)
	f(s)
	assert.Equal(t, []int{}, s)
	f2(&s)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, s)
}
