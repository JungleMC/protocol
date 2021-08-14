package protocol

import (
	"bufio"
	"golang.org/x/text/encoding/unicode"
)

// Legacy ping values
const pingHostText = "MC|PingHost"
const remainingBytes = 7
const LegacyPingPayload = 0xFE

var (
	utfEncoder = unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM).NewEncoder()
	utfDecoder = unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM).NewDecoder()
)

type HandshakePacket struct {
	ProtocolVersion int32
	ServerHost      string
	ServerPort      uint16
	NextState       int32
}

func (p *HandshakePacket) Id() int32 {
	return 0x00
}

func (p *HandshakePacket) Encode() []byte {
	buf := WriteVarInt32(p.ProtocolVersion)
	buf = append(buf, WriteString(p.ServerHost)...)
	buf = append(buf, WriteUint16(p.ServerPort)...)
	return append(buf, WriteVarInt32(p.NextState)...)
}

func (p *HandshakePacket) Decode(buf *bufio.Reader) {
	p.ProtocolVersion = ReadVarInt32(buf)
	p.ServerHost = ReadString(buf)
	p.ServerPort = ReadUint16(buf)
	p.NextState = ReadVarInt32(buf)
}

type LegacyPingPacket struct {
	Payload         byte
	ProtocolVersion byte
	PluginMessage   byte
	Hostname        string
	Port            int32
}

func (p *LegacyPingPacket) Id() int32 {
	return 0xFE
}

func (p *LegacyPingPacket) Encode() []byte {
	buf := make([]byte, 0)
	buf = append(buf, p.Payload)
	buf = append(buf, p.ProtocolVersion)
	buf = append(buf, p.PluginMessage)

	pingHostBytes, _ := utfEncoder.Bytes([]byte(pingHostText))
	buf = append(buf, WriteInt16(int16(len(pingHostBytes)))...)
	buf = append(buf, pingHostBytes...)

	hostnameBytes, _ := utfEncoder.Bytes([]byte(p.Hostname))
	buf = append(buf, WriteInt16(int16(remainingBytes+len(hostnameBytes)))...)
	buf = append(buf, hostnameBytes...)

	return append(buf, WriteInt32(p.Port)...)
}

func (p *LegacyPingPacket) Decode(buf *bufio.Reader) {
	p.Payload, _ = buf.ReadByte()
	p.ProtocolVersion, _ = buf.ReadByte()
	p.PluginMessage, _ = buf.ReadByte()

	pingHostLen := ReadUint16(buf)
	pingHost := make([]byte, pingHostLen)
	buf.Read(pingHost)

	remaining := ReadInt16(buf)
	hostnameBytes := make([]byte, remaining-remainingBytes)
	buf.Read(hostnameBytes)
	hostnameBytes, _ = utfDecoder.Bytes(hostnameBytes)
	p.Hostname = string(hostnameBytes)
	p.Port = ReadInt32(buf)
}
