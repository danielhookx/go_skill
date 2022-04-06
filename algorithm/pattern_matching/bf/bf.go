package bf

func PatternMatching(target, pattern string) int {
	for i := 0; i < len(target); i++ {
		p := i
		for j := 0; j < len(pattern); j++ {
			if pattern[j] != target[p] {
				break
			}
			p++
		}
		if p >= i+len(pattern) {
			return i
		}
	}
	return -1
}
