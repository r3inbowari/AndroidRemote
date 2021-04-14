package db

import (
	bilicoin "RemoteServer/utils"
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var ctx = context.Background()

var Rdb *redis.Client

// ExampleClient example
func ExampleClient() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     bilicoin.GetConfig().RdbURL,
		Password: bilicoin.GetConfig().RdbPassword,
		DB:       0,
	})

	err := Rdb.Set(context.TODO(), "key", "value", time.Second*10).Err()
	// err := Rdb.Set(context.TOD(), "key", "value", time.Second*time.Duration(bilicoin.GetConfig().DeviceRenewTTL)).Err()
	if err != nil {
		panic(err)
	}
}

// InitRDB 只使用 index 0 数据库
func InitRDB() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     bilicoin.GetConfig().RdbURL,
		Password: bilicoin.GetConfig().RdbPassword,
		DB:       bilicoin.GetConfig().RdbIndex,
	})
	err := Rdb.Set(context.TODO(), "run"+bilicoin.CreateMD5(time.Now().String()), "ok", time.Minute).Err()
	if err != nil {
		bilicoin.Fatal("[RDB] please check your redis configuration")
		os.Exit(12)
	}
	bilicoin.Info("[RDB] init redis client", logrus.Fields{"DBIndex": bilicoin.GetConfig().RdbIndex})
}
