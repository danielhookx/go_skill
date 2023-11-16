package sort

func QuickSort[T Orderliness](src []T) {
	quickSort(src, 0, len(src))
}

func quickSort[T Orderliness](src []T, lo, hi int) {
	if lo >= hi {
		return
	}
	pivot := partition(src, lo, hi)
	quickSort(src, lo, pivot)
	quickSort(src, pivot+1, hi)
}

func partition[T Orderliness](src []T, lo, hi int) int {
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
