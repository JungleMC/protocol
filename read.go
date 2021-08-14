package protocol

import (
	"bufio"
	"github.com/google/uuid"
	"math"
)

const stringMaxLength = 32767

func ReadBool(buf *bufio.Reader) bool {
	b, _ := buf.ReadByte()
	return b == 0x01
}

func ReadUint8(buf *bufio.Reader) uint8 {
	b, _ := buf.ReadByte()
	return b
}

func ReadUint16(buf *bufio.Reader) uint16 {
	b := make([]byte, 2)
	_, _ = buf.Read(b)
	return uint16(b[0])<<8 | uint16(b[1])
}

func ReadUint32(buf *bufio.Reader) uint32 {
	b := make([]byte, 4)
	_, _ = buf.Read(b)
	return uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
}

func ReadUint64(buf *bufio.Reader) uint64 {
	b := make([]byte, 8)
	_, _ = buf.Read(b)
	return uint64(b[0])<<56 | uint64(b[1])<<48 | uint64(b[2])<<40 | uint64(b[3])<<32 | uint64(b[4])<<24 | uint64(b[5])<<16 | uint64(b[6])<<8 | uint64(b[7])
}

func ReadInt8(buf *bufio.Reader) int8 {
	return int8(ReadUint8(buf))
}

func ReadInt16(buf *bufio.Reader) int16 {
	return int16(ReadUint16(buf))
}

func ReadInt32(buf *bufio.Reader) int32 {
	return int32(ReadUint32(buf))
}

func ReadInt64(buf *bufio.Reader) int64 {
	return int64(ReadUint64(buf))
}

func ReadFloat32(buf *bufio.Reader) float32 {
	return math.Float32frombits(ReadUint32(buf))
}

func ReadFloat64(buf *bufio.Reader) float64 {
	return math.Float64frombits(ReadUint64(buf))
}

func ReadString(buf *bufio.Reader) string {
	length := ReadVarInt32(buf)
	if length > stringMaxLength {
		panic("string exceeds max length")
	}
	data := make([]byte, length, length)
	_, _ = buf.Read(data)
	return string(data)
}

func ReadByteSlice(buf *bufio.Reader) []byte {
	v := make([]byte, ReadVarInt32(buf))
	buf.Read(v)
	return v
}

func ReadUUID(buf *bufio.Reader) uuid.UUID {
	v := make([]byte, 16)
	buf.Read(v)
	id := uuid.UUID{}
	copy(id[:], v)
	return id
}

func ReadVarInt32(buf *bufio.Reader) int32 {
	numRead := int32(0)
	result := int32(0)

	for {
		read, _ := buf.ReadByte()
		value := read & 0b01111111
		result |= int32(value) << (7 * numRead)
		numRead++
		if numRead > 5 {
			panic("varint32 is too big")
		}

		if (read & 0b10000000) == 0 {
			break
		}
	}
	return result
}

func ReadVarInt64(buf *bufio.Reader) int64 {
	numRead := int32(0)
	result := int64(0)

	for {
		read, _ := buf.ReadByte()
		value := read & 0b01111111
		result |= int64(value) << (7 * numRead)
		numRead++
		if numRead > 10 {
			panic("varint64 is too big")
		}

		if (read & 0b10000000) == 0 {
			break
		}
	}
	return result
}
