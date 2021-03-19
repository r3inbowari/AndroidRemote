package CloudGameServer

import (
	bilicoin "CloudGameServer/utils"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/sirupsen/logrus"
)

// scheme -> ws/ssl/tcp
// var scheme = "tcp"
// var host = "106.13.79.157"
// var port = "1883"

// ClientID 随机acm0-bjd2-fdi1-am81
// var ClientID = bson.NewObjectId().String()
// var Username = "r3inb"
// var Password = "159463"
// var base = "r3inbowari.top"

var client mqtt.Client = nil

func GetMQTTInstance() mqtt.Client {
	if client == nil || !client.IsConnectionOpen() {
		var err error
		client, err = processMQTTClient()
		if err != nil {
			bilicoin.Warn("mqtt broker connect failed")
		} else {
			bilicoin.Info("connected to mq", logrus.Fields{"host": *bilicoin.GetConfig().BrokerHost})
		}
	}
	return client
}

func processMQTTClient() (mqtt.Client, error) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(*bilicoin.GetConfig().BrokerScheme + "://" + *bilicoin.GetConfig().BrokerHost)
	if bilicoin.GetConfig().BrokerClientID != nil {
		opts.SetClientID(*bilicoin.GetConfig().BrokerClientID)
	}
	opts.SetUsername(*bilicoin.GetConfig().BrokerUsername)
	opts.SetPassword(*bilicoin.GetConfig().BrokerPassword)
	// opts.SetKeepAlive(2 * time.Second)
	// 默认消费方式
	//opts.SetDefaultPublishHandler(defaultPublishHandler)
	// ping超时
	//opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return c, nil
}

func MQTTMapping(topic string, callback mqtt.MessageHandler) bool {
	if token := GetMQTTInstance().Subscribe(topic, 1, callback); token.Wait() && token.Error() != nil {
		bilicoin.Warn("subscribe failed", logrus.Fields{"topic": topic})
		return false
	}
	bilicoin.Info("subscribed successfully", logrus.Fields{"topic": topic})
	return true
}

func MQTTPublish(topic string, payload interface{}) {
	token := GetMQTTInstance().Publish(topic, 1, false, payload)
	token.Wait()
}
