package protocol

import (
	"bufio"
	"github.com/google/uuid"
)

type DisconnectPacket struct {
	Reason string // TODO: chat message json
}

func (p *DisconnectPacket) Id() int32 {
	return 0x00
}

func (p *DisconnectPacket) Encode() []byte {
	return WriteString(p.Reason)
}

func (p *DisconnectPacket) Decode(buf *bufio.Reader) {
	p.Reason = ReadString(buf)
}

type EncryptionRequest struct {
	ServerId    string
	PublicKey   []byte
	VerifyToken []byte
}

func (p *EncryptionRequest) Id() int32 {
	return 0x01
}

func (p *EncryptionRequest) Encode() []byte {
	buf := WriteString(p.ServerId)
	buf = append(buf, WriteByteSlice(p.PublicKey)...)
	return append(buf, WriteByteSlice(p.VerifyToken)...)
}

func (p *EncryptionRequest) Decode(buf *bufio.Reader) {
	p.ServerId = ReadString(buf)
	p.PublicKey = ReadByteSlice(buf)
	p.VerifyToken = ReadByteSlice(buf)
}

type LoginSuccess struct {
	UUID     uuid.UUID
	Username string
}

func (p *LoginSuccess) Id() int32 {
	return 0x02
}

func (p *LoginSuccess) Encode() []byte {
	return append(p.UUID[:], WriteString(p.Username)...)
}

func (p *LoginSuccess) Decode(buf *bufio.Reader) {
	p.UUID = ReadUUID(buf)
	p.Username = ReadString(buf)
}

type SetCompressionPacket struct {
	Threshold int32
}

func (p *SetCompressionPacket) Id() int32 {
	return 0x03
}

func (p *SetCompressionPacket) Encode() []byte {
	return WriteVarInt32(p.Threshold)
}

func (p *SetCompressionPacket) Decode(buf *bufio.Reader) {
	p.Threshold = ReadVarInt32(buf)
}

type LoginPluginRequest struct {
	MessageId int32
	Channel   string
	Data      []byte
}

func (p *LoginPluginRequest) Id() int32 {
	return 0x04
}

func (p *LoginPluginRequest) Encode() []byte {
	buf := WriteVarInt32(p.MessageId)
	buf = append(buf, WriteString(p.Channel)...)
	return append(buf, p.Data...)
}

func (p *LoginPluginRequest) Decode(buf *bufio.Reader) {
	p.MessageId = ReadInt32(buf)
	p.Channel = ReadString(buf)
	p.Data = make([]byte, buf.Buffered())
	buf.Read(p.Data)
}
