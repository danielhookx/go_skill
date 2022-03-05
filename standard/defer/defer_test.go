package _defer

import "testing"

func Test_anonymous2(t *testing.T) {
	t.Logf("anonymousClosure: %d", anonymousClosure())
	t.Logf("namedClosure: %d", namedClosure())
	t.Logf("namedClosure2: %d", namedClosure2())
	t.Logf("namedNotClosure: %d", namedNotClosure())
}
