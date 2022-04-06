package string

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}

func reverse(x int) int {
	res := 0
	for x != 0 {
		res = res*10 + x%10
		x = x / 10
	}
	if res > 1<<31-1 || res < -1<<31 {
		return 0
	}
	return res
}

func firstUniqChar(s string) int {
	m := make(map[int32]int)
	for i, c := range s {
		if _, ok := m[c]; ok {
			m[c] = -i
		} else {
			m[c] = i
		}
	}
	min := 1<<31 - 1
	for _, v := range m {
		if v < 0 {
			continue
		}
		if v < min {
			min = v
		}
	}
	if min == 1<<31-1 {
		return -1
	}
	return min
}

//func strStr(haystack string, needle string) int {
//
//}

func kmp(haystack string, needle string) int {
	table := prefixTable(needle)

	var i, j = 0, 0
	for len(haystack)-i >= len(needle)-j && j < len(needle) {
		if haystack[i] == needle[j] {
			i++
			j++
		} else {
			if j == 0 {
				i++
				continue
			}
			j = table[j-1]
		}
	}
	if j >= len(needle) {
		return i - j
	}
	return -1
}

func prefixTable(pattern string) []int {
	if len(pattern) < 1 {
		return nil
	}
	table := make([]int, len(pattern))
	table[0] = 0
	for i, j := 1, 0; i < len(pattern); {
		//compare
		if pattern[j] == pattern[i] {
			table[i] = j + 1
			i++
			j++
		} else {
			if j-1 < 0 {
				table[i] = 0
				i++
				continue
			}
			j = table[j-1]
		}
	}
	return table
}
