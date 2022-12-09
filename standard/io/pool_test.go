package io

import (
	"testing"
)

func TestNewPool(t *testing.T) {
	p := NewPool(10, 1)
	b := p.free
	i := 0
	for b != nil {
		i++
		t.Logf("num:%d, ptr is: %p\n", i, b.buf)
		b = b.next
	}
}
