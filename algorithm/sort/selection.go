package sort

func SelectionSort(src []int) {
	// i 表示无序空间的第一个元素位置
	for i := 0; i < len(src); i++ {
		for j := i; j < len(src); j++ {
			if src[i] > src[j] {
				//swap
				src[i], src[j] = src[j], src[i]
			}
		}
	}
}
