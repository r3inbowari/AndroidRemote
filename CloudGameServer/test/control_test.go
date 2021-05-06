package test

import (
	"CloudGameServer/db"
	"CloudGameServer/device"
	"testing"
)

func TestInit(t *testing.T) {
	db.InitRDB()
	device.InitDeviceDispatcher()
}
