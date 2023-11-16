package sort

// min(O)=max(O)=avg(O): O(n^2)
func BubbleSort[T Orderliness](src []T) {
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

// min(O)= 1 max(O)=O(n^2)
func BubbleSortFlag[T Orderliness](src []T) {
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
