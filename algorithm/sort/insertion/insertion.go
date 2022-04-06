package insertion

func InsertionDesc(src []int) {
	for i := 0; i < len(src); i++ {
		for j := i; j-1 >= 0; j-- {
			if src[j] > src[j-1] {
				src[j], src[j-1] = src[j-1], src[j]
			}
		}
	}
}

func InsertionAsc(src []int) {
	for i := 0; i < len(src); i++ {
		for j := i; j-1 >= 0; j-- {
			if src[j] < src[j-1] {
				src[j], src[j-1] = src[j-1], src[j]
			}
		}
	}
}
