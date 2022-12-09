package sort

func BubbleSort(src []int) {
	// i 表示正在冒第几个泡泡
	for i := 1; i <= len(src); i++ {
		for j := 0; j < len(src)-i; j++ {
			if src[j] > src[j+1] {
				//swap
				src[j], src[j+1] = src[j+1], src[j]
			}
		}
	}
}

func BubbleSortFlag(src []int) {
	sorted := true
	// i 表示正在冒第几个泡泡
	for i := 1; i <= len(src); i++ {
		sorted = true
		for j := 0; j < len(src)-i; j++ {
			if src[j] > src[j+1] {
				//swap
				src[j], src[j+1] = src[j+1], src[j]
				// 但凡有交换，则可能无序
				sorted = false
			}
		}
		if sorted {
			return
		}
	}
}
