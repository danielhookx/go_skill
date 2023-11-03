package rune

import (
	"fmt"
	"testing"
)

func TestRune(t *testing.T) {
	s1 := "s"
	s2 := "s鷥"
	fmt.Printf("rune size:%d;  byte size:%d\n", len([]rune(s1)), len(s1))
	fmt.Printf("rune size:%d;  byte size:%d\n", len([]rune(s2)), len(s2))
	fmt.Printf("%x\n", s1)
	fmt.Printf("%x\n", s2)
	fmt.Printf("%U\n", 's')
	fmt.Printf("%U\n", '鷥')
	// 0x00-0xff
}
