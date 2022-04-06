package selection

func SelectionDesc(src []int) {
	for i := 0; i < len(src); i++ {
		for j := i; j < len(src); j++ {
			if src[i] < src[j] {
				src[i], src[j] = src[j], src[i]
			}
		}
	}
}

func SelectionAsc(src []int) {
	for i := 0; i < len(src); i++ {
		for j := i; j < len(src); j++ {
			if src[i] > src[j] {
				src[i], src[j] = src[j], src[i]
			}
		}
	}
}
