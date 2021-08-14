package protocol

import "bufio"

type StatusResponse struct {
	Response string
}

func (p *StatusResponse) Id() int32 {
	return 0x00
}

func (p *StatusResponse) Encode() []byte {
	return WriteString(p.Response)
}

func (p *StatusResponse) Decode(buf *bufio.Reader) {
	p.Response = ReadString(buf)
}

type StatusPong struct {
	Time int64
}

func (p *StatusPong) Id() int32 {
	return 0x01
}

func (p *StatusPong) Encode() []byte {
	return WriteInt64(p.Time)
}

func (p *StatusPong) Decode(buf *bufio.Reader) {
	p.Time = ReadInt64(buf)
}
