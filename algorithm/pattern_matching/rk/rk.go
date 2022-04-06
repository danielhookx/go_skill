package rk

var digitTable []int

func tableInit(len int) {
	if len < 1 {
		panic("tableInit length less than 1")
	}
	digitTable = make([]int, len)
	digitTable[0] = 1
	for i := 1; i < len; i++ {
		digitTable[i] = digitTable[i-1] * 26
	}
}

func RK(target, pattern string) int {
	//get pattern hash
	tableInit(len(pattern))
	patternHash := simpleHash(pattern)

	//
	if len(target) < len(pattern) {
		return -1
	}
	slideHash := simpleHash(target[:len(pattern)])
	if slideHash == patternHash {
		return 0
	}
	m := len(pattern)

	for i := 1; i <= len(target)-len(pattern); i++ {
		slideHash = 26*(slideHash-int(target[i-1]-'a')*digitTable[m-1]) + int(target[i+m-1]-'a')
		if slideHash == patternHash {
			return i
		}
	}
	return -1
}

func simpleHash(str string) int {
	hash := 0
	for i := 0; i < len(str); i++ {
		hash += int(str[i]-'a') * digitTable[len(str)-i-1]
	}
	return hash
}
