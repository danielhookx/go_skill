package bubble

//
func BubbleSortDesc(src []int) {
	for i := 0; i < len(src); i++ {
		for j := 0; j < len(src)-i-1; j++ {
			if src[j] < src[j+1] {
				src[j], src[j+1] = src[j+1], src[j]
			}
		}
	}
}

func BubbleSortAsc(src []int) {
	for i := 0; i < len(src); i++ {
		for j := 0; j < len(src)-i-1; j++ {
			if src[j] > src[j+1] {
				src[j], src[j+1] = src[j+1], src[j]
			}
		}
	}
}
