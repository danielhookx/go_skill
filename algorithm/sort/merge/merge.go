package merge

func MergeSort(src []int) {
	if len(src) <= 1 {
		return
	}
	MergeSort(src[:len(src)/2])
	MergeSort(src[len(src)/2:])
	merge(src, src[:len(src)/2], src[len(src)/2:])
}

func merge(src, before, after []int) {
	bf, af := 0, 0
	index := 0
	dst := make([]int, len(before)+len(after))
	for af < len(after) && bf < len(before) {
		if before[bf] > after[af] {
			dst[index] = after[af]
			af++
			index++
		} else {
			dst[index] = before[bf]
			bf++
			index++
		}
	}
	if bf < len(before) {
		copy(dst[index:], before[bf:])
	} else if af < len(after) {
		copy(dst[index:], after[af:])
	}
	copy(src[:], dst[:])
}
