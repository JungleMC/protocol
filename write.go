package protocol

import (
	"github.com/google/uuid"
	"math"
)

func WriteBool(v bool) []byte {
	if v {
		return []byte{0x01}
	}
	return []byte{0x00}
}

func WriteUint8(v uint8) []byte {
	return []byte{v}
}

func WriteUint16(v uint16) []byte {
	return []byte{byte(v >> 8), byte(v)}
}

func WriteUint32(v uint32) []byte {
	return []byte{byte(v >> 24), byte(v >> 16), byte(v >> 8), byte(v)}
}

func WriteUint64(v uint64) []byte {
	return []byte{byte(v >> 56), byte(v >> 48), byte(v >> 40), byte(v >> 32), byte(v >> 24), byte(v >> 16), byte(v >> 8), byte(v)}
}

func WriteInt8(v int8) []byte {
	return WriteUint8(uint8(v))
}

func WriteInt16(v int16) []byte {
	return WriteUint16(uint16(v))
}

func WriteInt32(v int32) []byte {
	return WriteUint32(uint32(v))
}

func WriteInt64(v int64) []byte {
	return WriteUint64(uint64(v))
}

func WriteFloat32(v float32) []byte {
	return WriteUint32(math.Float32bits(v))
}

func WriteFloat64(v float64) []byte {
	return WriteUint64(math.Float64bits(v))
}

func WriteString(v string) []byte {
	if len(v) > stringMaxLength {
		panic("string exceeds max length")
	}
	data := WriteVarInt32(int32(len(v)))
	return append(data, []byte(v)...)
}

func WriteByteSlice(v []byte) []byte {
	return append(WriteVarInt32(int32(len(v))), v...)
}

func WriteUUID(v uuid.UUID) []byte {
	return v[:]
}

func WriteVarInt32(v int32) []byte {
	buf := make([]byte, 0)
	uvalue := uint32(v)

	for {
		temp := byte(uvalue & 0b01111111)
		uvalue >>= 7
		if uvalue != 0 {
			temp |= 0b10000000
		}
		buf = append(buf, temp)

		if uvalue == 0 {
			break
		}
	}
	return buf
}

func WriteVarInt64(v int64) []byte {
	buf := make([]byte, 0)
	uvalue := uint64(v)
	for {
		temp := byte(uvalue & 0b01111111)
		uvalue >>= 7
		if uvalue != 0 {
			temp |= 0b10000000
		}
		buf = append(buf, temp)
		if uvalue == 0 {
			break
		}
	}
	return buf
}
