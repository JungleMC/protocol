package protocol

import "bufio"

type LoginStartPacket struct {
	Username string
}

func (p *LoginStartPacket) Id() int32 {
	return 0x00
}

func (p *LoginStartPacket) Encode() []byte {
	return WriteString(p.Username)
}

func (p *LoginStartPacket) Decode(buf *bufio.Reader) {
	p.Username = ReadString(buf)
}

type EncryptionResponsePacket struct {
	SharedSecret []byte
	VerifyToken  []byte
}

func (p *EncryptionResponsePacket) Id() int32 {
	return 0x01
}

func (p *EncryptionResponsePacket) Encode() []byte {
	return append(
		WriteByteSlice(p.VerifyToken),
		WriteByteSlice(p.SharedSecret)...,
	)
}

func (p *EncryptionResponsePacket) Decode(buf *bufio.Reader) {
	p.VerifyToken = ReadByteSlice(buf)
	p.SharedSecret = ReadByteSlice(buf)
}

type LoginPluginResponse struct {
	MessageID  int32
	Successful bool
	Data       []byte
}

func (p *LoginPluginResponse) Id() int32 {
	return 0x02
}

func (p *LoginPluginResponse) Encode() []byte {
	buf := WriteVarInt32(p.MessageID)
	buf = append(buf, WriteBool(p.Successful)...)
	if p.Successful {
		return append(buf, p.Data...)
	}
	return buf
}

func (p *LoginPluginResponse) Decode(buf *bufio.Reader) {
	p.MessageID = ReadVarInt32(buf)
	p.Successful = ReadBool(buf)
	if p.Successful {
		p.Data = make([]byte, buf.Buffered())
		buf.Read(p.Data)
	}
}
