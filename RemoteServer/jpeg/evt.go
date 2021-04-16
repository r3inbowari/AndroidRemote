package pics

import (
	"RemoteServer/event"
	"encoding/binary"
)

// ParseEvent
// Type:    int(bs[6] & 0x0F),
// X:       int(binary.BigEndian.Uint16(bs[7:9])),
// Y:       int(binary.BigEndian.Uint16(bs[9:11])),
// Contact: int((bs[6] & 0xF0) >> 4),
// Ts:      int64(binary.BigEndian.Uint64(bs[11:19])),
func ParseEvent(bs []byte) *event.EventResponse {
	if len(bs) != 12 {
		return nil
	}
	t := event.EventResponse{
		Type:    event.EventResponse_EventType(int(bs[3] & 0x0F)),
		Contact: int32((bs[3] & 0xF0) >> 4),
		X:       int32(binary.BigEndian.Uint16(bs[4:6])),
		Y:       int32(binary.BigEndian.Uint16(bs[6:8])),
		Ts:      int64(binary.BigEndian.Uint32(bs[8:])),
	}
	return &t
}

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
