package _defer

func anonymousClosure() int {
	var a = 0

	defer func() {
		a++
	}()

	return a
}

func namedClosure() (a int) {
	defer func() {
		a++
	}()

	return a
}

func namedClosure2() (a int) {
	var b = 1

	defer func() {
		a = b
	}()
	b++
	return a
}

func namedNotClosure() (a int) {
	var b = 1

	defer func(b int) {
		a = b
	}(b)
	b++
	return a
}
