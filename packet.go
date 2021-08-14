package protocol

import (
	"bufio"
)

type Packet interface {
	Id() int32
	Encode() []byte
	Decode(buffer *bufio.Reader)
}

func NewPacket(state ConnectionState, direction Direction, id int32) Packet {
	if direction != ServerToClient {
		switch state {
		case Handshake:
			return clientboundHandshakePacket(id)
		case Status:
			return clientboundStatusPacket(id)
		case Login:
			return clientboundLoginPacket(id)
		case Play:
			return clientboundPlayPacket(id)
		}
	} else {
		switch state {
		case Handshake:
			return serverboundHandshakePacket(id)
		case Status:
			return serverboundStatusPacket(id)
		case Login:
			return serverboundLoginPacket(id)
		case Play:
			return serverboundPlayPacket(id)
		}
	}
	return nil
}

func clientboundHandshakePacket(_ int32) Packet {
	// no data
	return nil
}

func serverboundHandshakePacket(id int32) Packet {
	switch id {
	case 0x00:
		return &HandshakePacket{}
	case 0xFE:
		return &LegacyPingPacket{}
	}
	return nil
}

func clientboundStatusPacket(id int32) Packet {
	switch id {
	case 0x00:
		return &StatusResponsePacket{}
	case 0x01:
		return &StatusPongPacket{}
	}
	return nil
}

func serverboundStatusPacket(id int32) Packet {
	switch id {
	case 0x00:
		return &StatusRequestPacket{}
	case 0x01:
		return &StatusPingPacket{}
	}
	return nil
}

func clientboundLoginPacket(id int32) Packet {
	switch id {
	case 0x00:
		return &DisconnectPacket{}
	case 0x01:
		return &EncryptionRequest{}
	case 0x02:
		return &LoginSuccess{}
	case 0x03:
		return &SetCompressionPacket{}
	case 0x04:
		return &LoginPluginRequest{}
	}
	return nil
}

func serverboundLoginPacket(id int32) Packet {
	switch id {
	case 0x00:
		return &LoginStartPacket{}
	case 0x01:
		return &EncryptionResponsePacket{}
	case 0x02:
		return &LoginPluginResponse{}
	}
	return nil
}

func clientboundPlayPacket(id int32) Packet {
	return nil
}

func serverboundPlayPacket(id int32) Packet {
	return nil
}
