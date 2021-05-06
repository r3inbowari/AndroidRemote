package device

import (
	"CloudGameServer/db"
	"golang.org/x/net/context"
)

func InitDeviceDispatcher() {
	re := db.Rdb.Keys(context.TODO(), "reg:info:*")
	sl, _ := re.Result()
	println(len(sl))


	println(Exits("deab9dbaaa74541da"))


	db.Rdb.HSet(context.TODO(), "sessionMap", "deab9dbaaa74541da", "aaaa")

	result, err := db.Rdb.Get(context.TODO(), "deab9dbaaa74541da").Result()
	if err != nil {
		return
	}
	println(result)
}

func Exits(deviceID string) bool {
	res := db.Rdb.Exists(context.TODO(), "reg:info:"+deviceID)
	result, err := res.Result()
	if err != nil || result == 0 {
		return false
	}
	return true
}

func Attach() {

}

func Detach() {

}