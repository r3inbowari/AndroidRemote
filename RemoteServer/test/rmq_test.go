package test

import (
	"RemoteServer/queue"
	bilicoin "RemoteServer/utils"
	"fmt"
	"github.com/adjust/rmq/v3"
	"github.com/go-redis/redis/v7"
	"log"
	"testing"
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

	things, err := connection.OpenQueue("things")
	if err != nil {
		panic(err)
	}

	delivery := fmt.Sprintf("delivery %d", 1)

	if err = things.Publish(delivery); err != nil {
		log.Printf("failed to publish: %s", err)
	}
}
