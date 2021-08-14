package protocol

import "bufio"

type StatusRequest struct{}

func (p *StatusRequest) Id() int32 {
	return 0x00
}

func (p *StatusRequest) Encode() []byte {
	return []byte{}
}

func (p *StatusRequest) Decode(_ *bufio.Reader) {}

type StatusPing struct {
	Time int64
}

func (p *StatusPing) Id() int32 {
	return 0x01
}

func (p *StatusPing) Encode() []byte {
	return WriteInt64(p.Time)
}

func (p *StatusPing) Decode(buf *bufio.Reader) {
	p.Time = ReadInt64(buf)
}
