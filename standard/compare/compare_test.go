package compare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanCompare(t *testing.T) {
	//Integer,Float,string,Boolean,Complex,Pointer,Channel,Interface,Array
	ch1 := make(chan int)
	ch2 := ch1
	ch3 := make(chan struct{})
	ch4 := make(chan struct{})
	assert.True(t, ch1 == ch2)
	assert.False(t, ch3 == ch4)
	// Invalid operation: ch1 == ch3 (mismatched types chan int and chan struct{}
	// ch1 == ch3

}
