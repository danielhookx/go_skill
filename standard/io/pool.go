package io

import "sync"

type Buffer struct {
	buf  []byte
	next *Buffer
}

func (b *Buffer) Buffer() []byte {
	return b.buf
}

type Pool struct {
	lock   sync.Mutex
	free   *Buffer
	num    int
	size   int
	factor int
}

func NewPool(num, size int) *Pool {
	p := Pool{
		num:    num,
		size:   size,
		factor: num * size,
	}
	p.grow()
	return &p
}

func (p *Pool) grow() {
	buf := make([]byte, p.factor)
	bp := make([]Buffer, p.num)
	p.free = &bp[0]
	b := p.free
	for i := 1; i < p.num; i++ {
		b.buf = buf[(i-1)*p.size : i*p.size]
		b.next = &bp[i]
		b = b.next
	}
	b.buf = buf[(p.num-1)*p.size : p.num*p.size]
	b.next = nil
}

func (p *Pool) Put(b *Buffer) {
	p.lock.Lock()
	b.next = p.free
	p.free = b
	p.lock.Unlock()
}

func (p *Pool) Get() *Buffer {
	b := p.free
	if b == nil {
		p.grow()
		b = p.free
	}
	p.free = b.next
	return b
}
