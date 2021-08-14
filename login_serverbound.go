package protocol

import "bufio"

type LoginStart struct {
	Username string
}

func (p *LoginStart) Id() int32 {
	return 0x00
}

func (p *LoginStart) Encode() []byte {
	return WriteString(p.Username)
}

func (p *LoginStart) Decode(buf *bufio.Reader) {
	p.Username = ReadString(buf)
}

type EncryptionResponse struct {
	SharedSecret []byte
	VerifyToken  []byte
}

func (p *EncryptionResponse) Id() int32 {
	return 0x01
}

func (p *EncryptionResponse) Encode() []byte {
	return append(
		WriteByteSlice(p.SharedSecret),
		WriteByteSlice(p.VerifyToken)...,
	)
}

func (p *EncryptionResponse) Decode(buf *bufio.Reader) {
	p.SharedSecret = ReadByteSlice(buf)
	p.VerifyToken = ReadByteSlice(buf)
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
