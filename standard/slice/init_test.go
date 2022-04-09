package slice

import "testing"

func Test_InitSlice(t *testing.T) {
	// Using slice literal:
	s1 := []int{1, 2, 3, 4, 5}

	//Using an Array:
	a1 := [5]int{1, 2, 3, 4, 5}
	s2 := a1[:len(a1)-1]

	//Using already Existing Slice:
	s3 := s2[:]

	//Using make() function:
	s4 := make([]int, 4, 7)
	s5 := make([]int, 7)

	t.Logf("s1:val=%v, len=%d, cap=%d \n", s1, len(s1), cap(s1))
	t.Logf("s2:val=%v, len=%d, cap=%d \n", s2, len(s2), cap(s2))
	t.Logf("s3:val=%v, len=%d, cap=%d \n", s3, len(s3), cap(s3))
	t.Logf("s4:val=%v, len=%d, cap=%d \n", s4, len(s4), cap(s4))
	t.Logf("s5:val=%v, len=%d, cap=%d \n", s5, len(s5), cap(s5))
}

func Test_SliceAppend(t *testing.T) {
	s := make([]int, 3, 3)
	s1 := append(s, []int{1, 2, 3, 4}...)
	// the result is
	//s1:len=7, cap=8
	//new_cap = cap(64 / 8)
	// 7 * 8 byte = 56B
	//span level 6 = 64B
	t.Logf("s1:len=%d, cap=%d \n", len(s1), cap(s1))
	t.Logf("s:ptr=%p, s1:ptr=%p \n", s, s1)

	s = make([]int, 256, 256)
	s1 = append(s, 1)
	// the result is
	//s1:len=257, cap=512
	//new_cap = double old_cap
	t.Logf("s1:len=%d, cap=%d \n", len(s1), cap(s1))

	s = make([]int, 1024, 1024)
	s1 = append(s, 1)
	// the result is
	//s1:len=1025, cap=1280
	//new_cap = old_cap * 1.25
	t.Logf("s1:len=%d, cap=%d \n", len(s1), cap(s1))
}

func Benchmark_SliceCopy(b *testing.B) {
	benchmarks := []struct {
		name string
		dst  []int
		src  []int
	}{
		{
			name: "zero",
			dst:  make([]int, 0),
			src:  make([]int, 0),
		}, {
			name: "one",
			dst:  make([]int, 1),
			src:  make([]int, 1),
		}, {
			name: "two",
			dst:  make([]int, 2),
			src:  make([]int, 2),
		}, {
			name: "more",
			dst:  make([]int, 100),
			src:  make([]int, 100),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				copy(bm.dst, bm.src)
			}
		})
	}
	/*
		Benchmark_SliceCopy
		Benchmark_SliceCopy/zero
		Benchmark_SliceCopy/zero-10         	1000000000	         0.5941 ns/op
		Benchmark_SliceCopy/one
		Benchmark_SliceCopy/one-10          	645597904	         1.851 ns/op
		Benchmark_SliceCopy/two
		Benchmark_SliceCopy/two-10          	622364944	         1.915 ns/op
		Benchmark_SliceCopy/more
		Benchmark_SliceCopy/more-10         	96457857	        11.84 ns/op
		PASS
	*/
}

/*
type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

*/
func Test_TheSlicePointTo(t *testing.T) {
	// Init the array that the slice point to
	a := [5]int64{1, 2, 3, 4, 5}
	t.Logf("%p", &a)
	s := a[1:4]
	t.Logf("%p", s)
	t.Logf("%p", &s[len(s)-1])
	t.Logf("%d, %d", len(s), cap(s))
}
