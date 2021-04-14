package control

import (
	"testing"
)

func TestDelSession(t *testing.T) {
	CancelSession("un exist")
}
