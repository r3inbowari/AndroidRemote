package event

import (
	"encoding/binary"
)

// test func

func GetSessionID(bs []byte) string {
	if len(bs) != 51 {
		return ""
	}
	return string(bs[19:])
}

func GetContact(bs []byte) int {
	if len(bs) != 51 {
		return 0
	}
	return int((bs[6] & 0xF0) >> 4)
}

func GetType(bs []byte) int {
	if len(bs) != 51 {
		return 0
	}
	return int(bs[6] & 0x0F)
}

func GetX(bs []byte) int {
	if len(bs) != 51 {
		return 0
	}
	return int(binary.BigEndian.Uint16(bs[7:9]))
}

func GetY(bs []byte) int {
	if len(bs) != 51 {
		return 0
	}
	return int(binary.BigEndian.Uint16(bs[9:11]))
}

func GetTs(bs []byte) int64 {
	if len(bs) != 51 {
		return 0
	}
	return int64(binary.BigEndian.Uint64(bs[11:19]))
}

func GetLen(bs []byte) int {
	if len(bs) != 51 {
		return 0
	}
	return int(bs[5])
}
