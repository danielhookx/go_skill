package io

import (
	"bufio"
	"math/rand"
	"strings"
	"testing"
	"time"

	xbufio "github.com/oofpgDLD/go_skill/standard/bufio"
)

func gen1KBStr() string {
	maxLen := 1024
	char := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rlt := make([]uint8, maxLen)
	rand.Seed(time.Now().Unix())
	for i := 0; i < maxLen; i++ {
		rlt[i] = char[rand.Intn(len(char))]
	}
	return string(rlt)
}

func Test_gen1KBStr(t *testing.T) {
	t.Log(gen1KBStr())
}

func Benchmark_IOString(b *testing.B) {
	str := gen1KBStr()
	for i := 0; i < b.N; i++ {
		strings.NewReader(str)
	}
}

func Benchmark_IOBuf(b *testing.B) {
	str := gen1KBStr()
	for i := 0; i < b.N; i++ {
		bufio.NewReaderSize(strings.NewReader(str), 1024)
	}
}

func Benchmark_IOBufPool(b *testing.B) {
	p := NewPool(10000, 1024)
	str := gen1KBStr()
	for i := 0; i < b.N; i++ {
		var rr xbufio.Reader
		bp := p.Get()
		rr.ResetBuf(strings.NewReader(str), bp.Buffer())
		p.Put(bp)
	}
}
