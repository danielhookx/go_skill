package _defer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnonymousClosure(t *testing.T) {
	f := func() int {
		var a = 0

		defer func() {
			a++
		}()

		return a
	}
	assert.Equal(t, 0, f())
}

func TestNamedClosure(t *testing.T) {
	f := func() (a int) {
		defer func() {
			a++
		}()
		return a
	}
	assert.Equal(t, 1, f())
}

func TestNamedClosure2(t *testing.T) {
	f := func() (a int) {
		b := 1

		defer func() {
			a = b
		}()
		b++
		return a
	}
	assert.Equal(t, 2, f())
}

func TestNamedNotClosure(t *testing.T) {
	f := func() (a int) {
		b := 1

		defer func(b int) {
			a = b
		}(b)
		return a
	}
	assert.Equal(t, 1, f())
}
