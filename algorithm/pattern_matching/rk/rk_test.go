package rk

import "testing"

func Test_tableInit(t *testing.T) {
	tableInit(3)
	for _, v := range digitTable {
		t.Log(v)
	}
}

func TestRK(t *testing.T) {
	t.Log(RK("hello", "ll"))
	t.Log(RK("he", "e"))
	t.Log(RK("he", "he"))
}
