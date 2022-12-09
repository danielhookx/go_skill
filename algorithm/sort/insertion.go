package sort

func InsertionSort(src []int) {
	// i 表示无序空间的第一个元素位置
	for i := 0; i < len(src); i++ {
		// 从有序区间的尾部插入
		for j := i; j > 0; j-- {
			if src[j] < src[j-1] {
				// swap
				src[j], src[j-1] = src[j-1], src[j]
			}
		}
	}
}
