package device

import (
	"RemoteServer/db"
	bilicoin "RemoteServer/utils"
	"golang.org/x/net/context"
	"time"
)

type DInfo struct {
}

// Renew the renewal of the device reg, the information of devices will has 120 seconds delay.
func Renew(deviceID, deviceInfo string) error {
	return db.Rdb.Set(context.TODO(), "reg:info:"+deviceID, deviceInfo, time.Second*time.Duration(bilicoin.GetConfig().DeviceRenewTTL)).Err()
}
