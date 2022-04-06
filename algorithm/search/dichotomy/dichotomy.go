package dichotomy

func recursively(src []int, lo, hi, search int) int {
	if lo > hi {
		return -1
	}
	mid := (lo + hi) / 2
	if src[mid] == search {
		return mid
	} else if src[mid] < search {
		return recursively(src, mid+1, hi, search)
	} else {
		return recursively(src, lo, mid-1, search)
	}
}

func Cycle(src []int, search int) int {
	lo := 0
	hi := len(src) - 1
	for lo <= hi {
		mid := (lo + hi) / 2
		if src[mid] == search {
			return mid
		} else if src[mid] < search {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}
