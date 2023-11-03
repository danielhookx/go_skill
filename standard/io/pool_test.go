package io

import (
	"bufio"
	"bytes"
	"io"
	"math/rand"
	"testing"
	"time"

	xbufio "github.com/oofpgDLD/go_skill/standard/bufio"
	"github.com/stretchr/testify/assert"
)

func gen1KB() []byte {
	maxLen := 1024
	char := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rlt := make([]uint8, maxLen)
	rand.Seed(time.Now().Unix())
	for i := 0; i < maxLen; i++ {
		rlt[i] = char[rand.Intn(len(char))]
	}
	return rlt
}

func Test_gen1KBStr(t *testing.T) {
	assert.Equal(t, 1024, len(gen1KB()))
}

var reader *bytes.Reader
var ioReader io.Reader
var xReader xbufio.Reader

func Benchmark_IOString(b *testing.B) {
	data := gen1KB()
	for i := 0; i < b.N; i++ {
		reader = bytes.NewReader(data)
		_ = reader
	}
}

func Benchmark_IOBuf(b *testing.B) {
	data := gen1KB()
	for i := 0; i < b.N; i++ {
		ioReader = bufio.NewReaderSize(bytes.NewReader(data), 1024)
		_ = ioReader
	}
}

func Benchmark_IOBufPool(b *testing.B) {
	data := gen1KB()
	p := NewPool(10000, 1024)
	for i := 0; i < b.N; i++ {
		bp := p.Get()
		xReader.ResetBuf(bytes.NewReader(data), bp.Buffer())
		_ = xReader
		p.Put(bp)
	}
}

func Benchmark_IOBufPool2(b *testing.B) {
	data := gen1KB()
	p := NewPool(10000, 1024)
	for i := 0; i < b.N; i++ {
		bp := p.Get()
		ioReader = io.TeeReader(bytes.NewReader(data), bytes.NewBuffer(bp.Buffer()))
		_ = ioReader
		p.Put(bp)
	}
}

func Benchmark_IOBufPool3(b *testing.B) {
	data := gen1KB()
	for i := 0; i < b.N; i++ {
		io.TeeReader(bytes.NewReader(data), GetBuffer())
	}
}
