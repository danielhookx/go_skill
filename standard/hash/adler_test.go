package hash

import (
	"hash/crc32"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	h1 := crc32.NewIEEE()
	h1.Write([]byte("a"))

	h2 := crc32.NewIEEE()
	h2.Write([]byte("b"))
	h2.Sum32()

	assert.NotEqual(t, h1.Sum32(), h2.Sum32())

	h1.Write([]byte("b"))
	h2.Write([]byte("a"))
	assert.Equal(t, h1.Sum32(), h2.Sum32())
}
