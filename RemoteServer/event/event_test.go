package event

import (
	"testing"
)

func TestEncode(t *testing.T) {
	var raw = []byte{
		0x45, 0x56, 0x45, 0x4E, 0X54,
		0x33,
		0x10,       // 0001 0000 -> Contact 1, Type,
		0x01, 0xF4, // 500 X
		0x01, 0xF4, // 500 Y
		0x00, 0x00, 0x00, 0x00, 0x60, 0x79, 0x86, 0xA6, // 1618577062 Ts
	}

	println(raw)

	s := "a8f5f167f44f4964e6c998dee827110c"
	ak := []byte(s)
	println(ak)

	raw = append(raw, ak...)

	println(raw)

	// get len
	println(GetLen(raw))
	println(GetTs(raw))
	println(GetX(raw))
	println(GetY(raw))
	println(GetContact(raw))
	println(GetType(raw))
	println(GetSessionID(raw))
}
