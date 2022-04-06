package quick

func QuickSort(src []int, lo, hi int) {
	if lo >= hi {
		return
	}
	pivot := partition(src, lo, hi)
	QuickSort(src, lo, pivot)
	QuickSort(src, pivot+1, hi)
}

func partition(src []int, lo, hi int) int {
	pivot := src[hi-1]
	i := lo

	for j := lo; j < hi; j++ {
		if src[j] < pivot {
			src[i], src[j] = src[j], src[i]
			i++
		}
	}
	src[i], src[hi-1] = src[hi-1], src[i]
	return i
}
