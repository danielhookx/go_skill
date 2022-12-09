package proto

import (
	"bufio"
	"encoding/binary"
	"errors"
)

const (
	// MaxBodySize max proto Body size
	MaxBodySize = int32(1 << 14)
)

const (
	// size
	_packSize      = 4
	_headerSize    = 2
	_verSize       = 2
	_rawHeaderSize = _packSize + _headerSize + _verSize
	_maxPackSize   = MaxBodySize + int32(_rawHeaderSize)
	// offset
	_packOffset   = 0
	_headerOffset = _packOffset + _packSize
	_verOffset    = _headerOffset + _headerSize
	_opOffset     = _verOffset + _verSize
)

var (
	// ErrProtoPackLen proto packet len error
	ErrProtoPackLen = errors.New("default server codec pack length error")
	// ErrProtoHeaderLen proto header len error
	ErrProtoHeaderLen = errors.New("default server codec header length error")
)

func Write(wr *bufio.Writer, p *Frame) error {
	packLen := _rawHeaderSize + uint32(len(p.Body))
	//
	buf := make([]byte, _rawHeaderSize)

	binary.BigEndian.PutUint32(buf[_packOffset:], packLen)
	binary.BigEndian.PutUint16(buf[_headerOffset:], uint16(_rawHeaderSize))
	binary.BigEndian.PutUint16(buf[_verOffset:], p.Ver)
	_, err := wr.Write(buf)
	if err != nil {
		return err
	}
	if p.Body != nil {
		_, err = wr.Write(p.Body)
	}
	return err
}

func Read(rr *bufio.Reader, p *Frame) error {
	rawHeader, err := pop(rr, _rawHeaderSize)
	if err != nil {
		return err
	}

	packLen := binary.BigEndian.Uint32(rawHeader[_packOffset:_headerOffset])
	headerLen := binary.BigEndian.Uint16(rawHeader[_headerOffset:_verOffset])
	p.Ver = binary.BigEndian.Uint16(rawHeader[_verOffset:_opOffset])
	if packLen > uint32(_maxPackSize) {
		return ErrProtoPackLen
	}
	if headerLen != _rawHeaderSize {
		return ErrProtoHeaderLen
	}
	if bodyLen := int(packLen - uint32(headerLen)); bodyLen > 0 {
		p.Body, err = pop(rr, bodyLen)
		if err != nil {
			return err
		}
	} else {
		p.Body = nil
	}
	return nil
}

func pop(rr *bufio.Reader, n int) ([]byte, error) {
	rawHeader, err := rr.Peek(n)
	if err != nil {
		return nil, err
	}
	discarded, err := rr.Discard(n)
	if err != nil {
		return nil, err
	}
	if discarded != n {
		return nil, errors.New("pop discarded len")
	}
	return rawHeader, nil
}
