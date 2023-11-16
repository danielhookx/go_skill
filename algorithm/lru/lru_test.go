package lru

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCache(t *testing.T) {
	c := NewCache[int, string](2)
	c.Add(1, "1")
	c.Add(2, "2")

	v, ok := c.Get(1)
	assert.Equal(t, true, ok)
	assert.EqualValues(t, "1", v)
	v, ok = c.Get(2)
	assert.Equal(t, true, ok)
	assert.EqualValues(t, "2", v)

	c.Add(3, "3")
	v, ok = c.Get(3)
	assert.Equal(t, true, ok)
	assert.EqualValues(t, "3", v)
	v, ok = c.Get(1)
	assert.Equal(t, false, ok)

	v, ok = c.Get(2)
	assert.Equal(t, true, ok)
	assert.EqualValues(t, "2", v)

	c.Add(4, "4")
	v, ok = c.Get(4)
	assert.Equal(t, true, ok)
	assert.EqualValues(t, "4", v)
	v, ok = c.Get(3)
	assert.Equal(t, false, ok)

	c.Add(2, "haha")
	c.Add(5, "5")
	v, ok = c.Get(2)
	assert.Equal(t, true, ok)
	assert.EqualValues(t, "haha", v)
	v, ok = c.Get(5)
	assert.Equal(t, true, ok)
	assert.EqualValues(t, "5", v)
}
