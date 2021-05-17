package test

import (
	"RemoteServer/queue"
	bilicoin "RemoteServer/utils"
	"encoding/json"
	"fmt"
	"github.com/adjust/rmq/v3"
	"github.com/go-redis/redis/v7"
	"log"
	"testing"
	"time"
)

func TestRMQ(t *testing.T) {
	go queue.Cleaner()
	select {}
}

func TestProduct(t *testing.T) {
	rInstance := redis.NewClient(&redis.Options{
		Addr:     bilicoin.GetConfig().RdbURL,
		Password: bilicoin.GetConfig().RdbPassword,
		DB:       1,
	})

	connection, err := rmq.OpenConnectionWithRedisClient("producer", rInstance, nil)
	if err != nil {
		panic(err)
	}

	things, err := connection.OpenQueue("order")
	if err != nil {
		panic(err)
	}

	delivery := fmt.Sprintf("delivery %d", 1)

	if err = things.Publish(delivery); err != nil {
		log.Printf("failed to publish: %s", err)
	}
}

func TestOpenSender(t *testing.T) {
	rInstance := redis.NewClient(&redis.Options{
		Addr:     bilicoin.GetConfig().RdbURL,
		Password: bilicoin.GetConfig().RdbPassword,
		DB:       1,
	})

	connection, err := rmq.OpenConnectionWithRedisClient("producer", rInstance, nil)
	if err != nil {
		panic(err)
	}

	order, err := connection.OpenQueue("order")
	if err != nil {
		panic(err)
	}

	op := queue.MQOrder{
		Id:        "deab9dbaaa74541d",
		Stub:      bilicoin.CreateMD5(time.Now().String()),
		Data:      "",
		Msg:       "order",
		Operation: bilicoin.REQ_START_SENDER,
	}

	dat, err := json.Marshal(op)
	if err != nil {
		bilicoin.Fatal("error marshal")
	}

	if err = order.Publish(string(dat)); err != nil {
		bilicoin.Fatal("failed to publish")
	}
}

func TestPauseSender(t *testing.T) {
	rInstance := redis.NewClient(&redis.Options{
		Addr:     bilicoin.GetConfig().RdbURL,
		Password: bilicoin.GetConfig().RdbPassword,
		DB:       1,
	})

	connection, err := rmq.OpenConnectionWithRedisClient("producer", rInstance, nil)
	if err != nil {
		panic(err)
	}

	order, err := connection.OpenQueue("order")
	if err != nil {
		panic(err)
	}

	op := queue.MQOrder{
		Id:        "deab9dbaaa74541d",
		Stub:      bilicoin.CreateMD5(time.Now().String()),
		Data:      "",
		Msg:       "order",
		Operation: bilicoin.REQ_PAUSE_SENDER,
	}

	dat, err := json.Marshal(op)
	if err != nil {
		bilicoin.Fatal("error marshal")
	}

	if err = order.Publish(string(dat)); err != nil {
		bilicoin.Fatal("failed to publish")
	}
}

func TestOpenApp(t *testing.T) {
	rInstance := redis.NewClient(&redis.Options{
		Addr:     bilicoin.GetConfig().RdbURL,
		Password: bilicoin.GetConfig().RdbPassword,
		DB:       1,
	})

	connection, err := rmq.OpenConnectionWithRedisClient("producer", rInstance, nil)
	if err != nil {
		panic(err)
	}

	order, err := connection.OpenQueue("order")
	if err != nil {
		panic(err)
	}

	op := queue.MQOrder{
		Id:        "deab9dbaaa74541d",
		Stub:      bilicoin.CreateMD5(time.Now().String()),
		Data:      "辐射避难所",
		Msg:       "order",
		Operation: bilicoin.REQ_OPEN_APP,
	}

	dat, err := json.Marshal(op)
	if err != nil {
		bilicoin.Fatal("error marshal")
	}

	if err = order.Publish(string(dat)); err != nil {
		bilicoin.Fatal("failed to publish")
	}
}

func TestCloseApp(t *testing.T) {
	rInstance := redis.NewClient(&redis.Options{
		Addr:     bilicoin.GetConfig().RdbURL,
		Password: bilicoin.GetConfig().RdbPassword,
		DB:       1,
	})

	connection, err := rmq.OpenConnectionWithRedisClient("producer", rInstance, nil)
	if err != nil {
		panic(err)
	}

	order, err := connection.OpenQueue("order")
	if err != nil {
		panic(err)
	}

	op := queue.MQOrder{
		Id:        "deab9dbaaa74541d",
		Stub:      bilicoin.CreateMD5(time.Now().String()),
		Data:      "com.xd.ro.roapk",
		Msg:       "order",
		Operation: bilicoin.REQ_CLOSE_APP,
	}

	dat, err := json.Marshal(op)
	if err != nil {
		bilicoin.Fatal("error marshal")
	}

	if err = order.Publish(string(dat)); err != nil {
		bilicoin.Fatal("failed to publish")
	}
}

//var AppMap sync.Map



// r3in.top
func TestOpenSenderRemote(t *testing.T) {
	rInstance := redis.NewClient(&redis.Options{
		Addr:     bilicoin.GetConfig().RdbURL,
		Password: bilicoin.GetConfig().RdbPassword,
		DB:       1,
	})

	connection, err := rmq.OpenConnectionWithRedisClient("producer", rInstance, nil)
	if err != nil {
		panic(err)
	}

	order, err := connection.OpenQueue("order")
	if err != nil {
		panic(err)
	}

	op := queue.MQOrder{
		Id:        "deab9dbaaa74541d",
		Stub:      bilicoin.CreateMD5(time.Now().String()),
		Data:      "",
		Msg:       "order",
		Operation: bilicoin.REQ_START_SENDER,
	}

	dat, err := json.Marshal(op)
	if err != nil {
		bilicoin.Fatal("error marshal")
	}

	if err = order.Publish(string(dat)); err != nil {
		bilicoin.Fatal("failed to publish")
	}
}

func TestOpenSenderSM(t *testing.T) {
	rInstance := redis.NewClient(&redis.Options{
		Addr:     bilicoin.GetConfig().RdbURL,
		Password: bilicoin.GetConfig().RdbPassword,
		DB:       1,
	})

	connection, err := rmq.OpenConnectionWithRedisClient("producer", rInstance, nil)
	if err != nil {
		panic(err)
	}

	order, err := connection.OpenQueue("order")
	if err != nil {
		panic(err)
	}

	op := queue.MQOrder{
		Id:        "e8d72afba1273295",
		Stub:      bilicoin.CreateMD5(time.Now().String()),
		Data:      "",
		Msg:       "order",
		Operation: bilicoin.REQ_START_SENDER,
	}

	dat, err := json.Marshal(op)
	if err != nil {
		bilicoin.Fatal("error marshal")
	}

	if err = order.Publish(string(dat)); err != nil {
		bilicoin.Fatal("failed to publish")
	}
}