package sort

func MergeSort[T Orderliness](src []T) {
	if len(src) < 2 {
		return
	}
	MergeSort(src[:len(src)/2])
	MergeSort(src[len(src)/2:])
	merge(src, len(src)/2)
}

func merge[T Orderliness](src []T, mid int) {
	merged := make([]T, len(src))
	left, right, current := 0, mid, 0
	for left < mid && right < len(src) {
		if src[left] < src[right] {
			merged[current] = src[left]
			left++
		} else {
			merged[current] = src[right]
			right++
		}
		current++
	}
	if left < mid {
		copy(merged[current:], src[left:mid])
	}
	if right < len(src) {
		copy(merged[current:], src[right:])
	}
	copy(src, merged)
}
