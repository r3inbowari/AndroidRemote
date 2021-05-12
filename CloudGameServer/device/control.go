package device

import (
	"CloudGameServer/db"
	"github.com/robfig/cron/v3"
	"golang.org/x/net/context"
	"sync"
)

func InitDeviceDispatcher() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("*/1 * * * * *", func() {
		getDeviceIDs()
	})
	c.Start()
}

var devicePool sync.Map

func Exits(deviceID string) bool {
	res := db.Rdb.Exists(context.TODO(), "reg:info:"+deviceID)
	result, err := res.Result()
	if err != nil || result == 0 {
		return false
	}
	return true
}

func PushSession2Redis(stub, device string) {
	db.Rdb.HSet(context.TODO(), "sessionMap", stub, device)
}

func DelSession4Redis(stub string) {
	db.Rdb.HDel(context.TODO(), "sessionMap", stub)
}

var Apply = make(chan string, 1)

//var dwg sync.WaitGroup

// Try2Apply 只允许一个请求进入临界区
func Try2Apply(stub string) bool {
	Apply <- stub
	// 获取一个可用容器
	deviceID, ok := GetAvailableContainer()
	if ok {
		// 交叉索引
		devicePool.Store(stub, deviceID)
		devicePool.Store(deviceID, stub)
		// 提交请求
		PushSession2Redis(stub, deviceID)
	}
	<-Apply
	return ok
}

func GetAvailableContainer() (string, bool) {
	res := getDeviceIDs()
	for _, v := range res {
		_, ok := devicePool.Load(v[9:])
		if !ok {
			return v[9:], true
		}
	}
	return "", false
}

var Apply2 = make(chan string, 1)

func Try2Detach(stub string) {
	Apply2 <- stub

	// 交叉索引清除
	deviceID, _ := devicePool.Load(stub)
	devicePool.Delete(stub)
	devicePool.Delete(deviceID)
	// 提交请求
	DelSession4Redis(stub)
	<-Apply2
}

var onlineDevice []string

func getDeviceIDs() []string {
	results, _ := db.Rdb.Keys(context.TODO(), "reg:info:*").Result()
	return results
}
