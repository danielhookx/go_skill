package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayInitialization(t *testing.T) {
	// Test initializing an array with values
	a1 := [5]int{1, 2, 3, 4, 5}
	a2 := [...]int{1, 2, 3, 4, 5}
	expected := [5]int{1, 2, 3, 4, 5}
	assert.Equal(t, expected, a1)
	assert.Equal(t, expected, a2)

	// Test initializing an array with default values
	var a3 [5]int
	expected2 := [5]int{0, 0, 0, 0, 0}
	assert.Equal(t, expected2, a3)
}
