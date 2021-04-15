package control

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"testing"
	"unsafe"
)

func TestDelSession(t *testing.T) {
	CancelSession("un exist")
}

func TestPtr(t *testing.T) {
	strPointerHex := fmt.Sprintf("%p", unsafe.Pointer(&redis.Client{}))
	println(strPointerHex)
}