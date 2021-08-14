package protocol

import "bufio"

type StatusRequestPacket struct{}

func (p *StatusRequestPacket) Id() int32 {
	return 0x00
}

func (p *StatusRequestPacket) Encode() []byte {
	return []byte{}
}

func (p *StatusRequestPacket) Decode(_ *bufio.Reader) {}

type StatusPingPacket struct {
	Time int64
}

func (p *StatusPingPacket) Id() int32 {
	return 0x01
}

func (p *StatusPingPacket) Encode() []byte {
	return WriteInt64(p.Time)
}

func (p *StatusPingPacket) Decode(buf *bufio.Reader) {
	p.Time = ReadInt64(buf)
}
