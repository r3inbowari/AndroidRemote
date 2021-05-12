package device

import (
	"CloudGameServer/db"
	"fmt"
	"github.com/robfig/cron/v3"
	"golang.org/x/net/context"
	"testing"
	"time"
)

func TestCron(t *testing.T) {
	c := cron.New(cron.WithSeconds())
	i := 1
	_, err := c.AddFunc("*/1 * * * * *", func() {
		fmt.Println("每分钟执行一次", i)
		i++
	})
	if err != nil {
		return
	}
	c.Start()

	time.Sleep(time.Minute * 5)
}

func TestAttach(t *testing.T) {
	db.InitRDB()
	Try2Apply("nihao")
	devicePool.Range(func(key, value interface{}) bool {
		println(key.(string))
		return true
	})

	Try2Apply("nihao1")
	devicePool.Range(func(key, value interface{}) bool {
		println(key.(string))
		return true
	})

}

func TestHashSet(t *testing.T) {
	db.InitRDB()

	println(Exits("deab9dbaaa74541da"))

	db.Rdb.HSet(context.TODO(), "sessionMap", "deab9dbaaa74541da", "aaaa")

	s := db.Rdb.HGet(context.TODO(), "sessionMap", "deab9dbaaa74541daa").Val()
	if s == "" {
		println("aaaaaaabbb")
	}
	println(s)
}

func TestHDel(t *testing.T) {
	db.InitRDB()

	Try2Detach("ad6d285f0e0fce2329e4290b8f4499ee")
}
