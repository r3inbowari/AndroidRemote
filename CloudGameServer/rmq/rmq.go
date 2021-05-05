package rmq

import (
	bilicoin "CloudGameServer/utils"
	"encoding/json"
	"fmt"
	"github.com/adjust/rmq/v3"
	"github.com/go-redis/redis/v7"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

const (
	prefetchLimit = 1000
	pollDuration  = 100 * time.Millisecond
	numConsumers  = 5

	reportBatchSize = 10000
	consumeDuration = time.Millisecond
	shouldLog       = false
)

func Cleaner() {
	// @warning must use go-redis/v7
	rInstance := redis.NewClient(&redis.Options{
		Addr:     bilicoin.GetConfig().RmqURL,
		Password: bilicoin.GetConfig().RmqPassword,
		DB:       bilicoin.GetConfig().RmqIndex,
	})
	connection, err := rmq.OpenConnectionWithRedisClient("cleaner", rInstance, nil)
	if err != nil {
		bilicoin.Fatal("[RMQ] connect failed", logrus.Fields{"err": err.Error()})
		os.Exit(14)
	}

	cleaner := rmq.NewCleaner(connection)
	bilicoin.Info("[RMQ] running -> garbage cleaner", logrus.Fields{"DEIndex": 1})

	for range time.Tick(time.Second) {
		cnt, err := cleaner.Clean()
		if err != nil {
			bilicoin.Fatal("[RMQ] failed to clean", logrus.Fields{"err": err.Error()})
			continue
		}
		if cnt > 0 {
			bilicoin.Info("[RMQ] cleaned", logrus.Fields{"count": int(cnt)})
		}
	}

}

var OrderQueue rmq.Queue

func InitRMQPub() {
	rInstance := redis.NewClient(&redis.Options{
		Addr:     bilicoin.GetConfig().RmqURL,
		Password: bilicoin.GetConfig().RmqPassword,
		DB:       bilicoin.GetConfig().RmqIndex,
	})

	bilicoin.Info("[RMQ] service init")

	connection, err := rmq.OpenConnectionWithRedisClient("producer", rInstance, nil)
	if err != nil {
		bilicoin.Fatal("[RMQ] connect failed, please check")
		panic(err)
	}

	OrderQueue, err = connection.OpenQueue("order")
	if err != nil {
		bilicoin.Fatal("[RMQ] run order failed, please check")
		panic(err)
	}
	bilicoin.Info("[RMQ] init rmq", logrus.Fields{"queueName": "order"})
}

func InitRMQ() {

	// @warning must use go-redis/v7
	rInstance := redis.NewClient(&redis.Options{
		Addr:     bilicoin.GetConfig().RmqURL,
		Password: bilicoin.GetConfig().RmqPassword,
		DB:       bilicoin.GetConfig().RmqIndex,
	})

	// 错误处理 channel
	errChan := make(chan error, 10)
	go logErrors(errChan)
	bilicoin.Info("[RMQ] running -> main service", logrus.Fields{"DBIndex": bilicoin.GetConfig().RmqIndex})
	go Cleaner()

	connection, err := rmq.OpenConnectionWithRedisClient("remote", rInstance, errChan)
	if err != nil {
		bilicoin.Fatal("[RMQ] connect failed", logrus.Fields{"err": err.Error()})
		os.Exit(14)
	}

	queue, err := connection.OpenQueue("order")
	if err != nil {
		bilicoin.Fatal("[RMQ] open rmq failed", logrus.Fields{"err": err.Error()})
		os.Exit(14)
	}

	if err := queue.StartConsuming(prefetchLimit, pollDuration); err != nil {
		bilicoin.Fatal("[RMQ] start consuming failed", logrus.Fields{"err": err.Error()})
		os.Exit(14)
	}

	//for i := 0; i < numConsumers; i++ {
	//	name := fmt.Sprintf("consumer %d", i)
	//	if _, err := rmq.AddConsumer(name, NewConsumer(i)); err != nil {
	//		bilicoin.Fatal("[RMQ] add consumer failed", logrus.Fields{"err": err.Error()})
	//	}
	//}
}

func NewConsumer(tag int) *Consumer {
	return &Consumer{
		name:   fmt.Sprintf("consumer%d", tag),
		count:  0,
		before: time.Now(),
	}
}

type Consumer struct {
	name   string
	count  int
	before time.Time
}

type MQOrder struct {
	Id        string      `json:"id"`
	Stub      string      `json:"stub"`
	Data      interface{} `json:"data"`
	Msg       string      `json:"msg"`
	Operation int         `json:"operation"`
}

type MQRsp struct {
	Id   string `json:"id"`   // 设备 ID
	Stub string `json:"stub"` // 操作存根
	Code int    `json:"code"` // 状态码
}

//func (consumer *Consumer) Consume(delivery rmq.Delivery) {
//	var entity MQOrder
//	if err := json.Unmarshal([]byte(delivery.Payload()), &entity); err != nil {
//		// handle json error
//		if err := delivery.Reject(); err != nil {
//			// handle reject error
//			bilicoin.Fatal("[RMQ] unknown error occurred during reject operation")
//		}
//		return
//	}
//
//	// 通过id执行执行开始投屏操作
//	if entity.Operation == bilicoin.REQ_START_SENDER {
//
//	}
//
//	switch entity.Operation {
//	case bilicoin.REQ_START_SENDER:
//		// @param type REQ_START_SENDER 1000
//		ret := control.GetSession(entity.Id)
//		if ret == nil {
//			// not exist session
//			bilicoin.Fatal("[RMQ] an offline or undefined device was accessed", logrus.Fields{"id": entity.Id, "stub": entity.Stub, "op": entity.Operation})
//			break
//		}
//		// send order to device
//		if err := ret.ChatStream.Send(&control.ChatResponse{Output: "order", Type: int32(entity.Operation)}); err != nil {
//			bilicoin.Info("[RMQ] device accessed", logrus.Fields{"id": entity.Id, "stub": entity.Stub, "op": entity.Operation})
//		}
//	case bilicoin.REQ_PAUSE_SENDER:
//		// @param type REQ_PAUSE_SENDER 1001
//		ret := control.GetSession(entity.Id)
//		if ret == nil {
//			// not exist session
//			bilicoin.Fatal("[RMQ] an offline or undefined device was accessed", logrus.Fields{"id": entity.Id, "stub": entity.Stub, "op": entity.Operation})
//			break
//		}
//		// send order to device
//		if err := ret.ChatStream.Send(&control.ChatResponse{Output: "order", Type: int32(entity.Operation)}); err != nil {
//			bilicoin.Info("[RMQ] device accessed", logrus.Fields{"id": entity.Id, "stub": entity.Stub, "op": entity.Operation})
//		}
//	case bilicoin.REQ_OPEN_APP:
//		// @param type REQ_OPEN_APP 1003
//		ret := control.GetSession(entity.Id)
//		if ret == nil {
//			// not exist session
//			bilicoin.Fatal("[RMQ] an offline or undefined device was accessed", logrus.Fields{"id": entity.Id, "stub": entity.Stub, "op": entity.Operation})
//			break
//		}
//		// send open order to device
//		if err := ret.ChatStream.Send(&control.ChatResponse{Output: entity.Data.(string), Type: int32(entity.Operation)}); err != nil {
//			bilicoin.Info("[RMQ] device accessed", logrus.Fields{"id": entity.Id, "stub": entity.Stub, "op": entity.Operation})
//		}
//
//	case bilicoin.REQ_CLOSE_APP:
//		// @param type REQ_CLOSE_APP 1004
//		ret := control.GetSession(entity.Id)
//		if ret == nil {
//			// not exist session
//			bilicoin.Fatal("[RMQ] an offline or undefined device was accessed", logrus.Fields{"id": entity.Id, "stub": entity.Stub, "op": entity.Operation})
//			break
//		}
//		// send open order to device
//		if err := ret.ChatStream.Send(&control.ChatResponse{Output: entity.Data.(string), Type: int32(entity.Operation)}); err != nil {
//			bilicoin.Info("[RMQ] device accessed", logrus.Fields{"id": entity.Id, "stub": entity.Stub, "op": entity.Operation})
//		}
//
//	default:
//		bilicoin.Fatal("[RMQ] illegal or unsupported device accessed", logrus.Fields{"id": entity.Id, "stub": entity.Stub, "op": entity.Operation})
//	}
//
//	// must ack, so you can return on the top switch-case
//	if err := delivery.Ack(); err != nil {
//		bilicoin.Fatal("[RMQ] unknown error occurred during ack operation", logrus.Fields{"id": entity.Id, "stub": entity.Stub, "op": entity.Operation})
//	}
//}

func logErrors(errChan <-chan error) {
	bilicoin.Info("[RMQ] running -> errors handler", logrus.Fields{"errorsCapacity": cap(errChan)})

	for err := range errChan {
		switch err := err.(type) {
		case *rmq.HeartbeatError:
			if err.Count == rmq.HeartbeatErrorLimit {
				bilicoin.Fatal("[RMQ] heartbeat error (limit): ", logrus.Fields{"err": err.Error()})
			} else {
				bilicoin.Fatal("[RMQ] heartbeat error: ", logrus.Fields{"err": err.Error()})
			}
		case *rmq.ConsumeError:
			bilicoin.Fatal("[RMQ] consume error: ", logrus.Fields{"err": err.Error()})
		case *rmq.DeliveryError:
			bilicoin.Fatal("[RMQ] delivery error: ", logrus.Fields{"err": err.Error(), "delivery": err.Delivery})
		default:
			bilicoin.Fatal("[RMQ] other error: ", logrus.Fields{"err": err.Error()})
		}
	}
}

func Order(id, stub, appName string, operation int) {
	op := MQOrder{
		Id:        id,
		Stub:      stub,
		Data:      appName,
		Msg:       "order",
		Operation: operation,
	}
	dat, err := json.Marshal(op)

	if err != nil {
		bilicoin.Fatal("error marshal")
	}

	if err = OrderQueue.Publish(string(dat)); err != nil {
		bilicoin.Fatal("failed to publish the order", logrus.Fields{"stub": stub})
	}
}
