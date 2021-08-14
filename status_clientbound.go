package protocol

import "bufio"

type StatusResponsePacket struct {
	Response string
}

func (p *StatusResponsePacket) Id() int32 {
	return 0x00
}

func (p *StatusResponsePacket) Encode() []byte {
	return WriteString(p.Response)
}

func (p *StatusResponsePacket) Decode(buf *bufio.Reader) {
	p.Response = ReadString(buf)
}

type StatusPongPacket struct {
	Time int64
}

func (p *StatusPongPacket) Id() int32 {
	return 0x01
}

func (p *StatusPongPacket) Encode() []byte {
	return WriteInt64(p.Time)
}

func (p *StatusPongPacket) Decode(buf *bufio.Reader) {
	p.Time = ReadInt64(buf)
}
