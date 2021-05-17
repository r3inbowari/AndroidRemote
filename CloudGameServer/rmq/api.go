package rmq

import (
	bilicoin "CloudGameServer/utils"
	"encoding/json"
	"github.com/sirupsen/logrus"
)

func OpenSender(stub, deviceID string) error {
	return wrapQueueRequest(stub, deviceID, "", "open sender", bilicoin.REQ_START_SENDER)
}

func CloseSender(stub, deviceID string) error {
	return wrapQueueRequest(stub, deviceID, "", "close sender", bilicoin.REQ_PAUSE_SENDER)
}

func OpenApp(stub, deviceID, appName string) error {
	return wrapQueueRequest(stub, deviceID, appName, "open app", bilicoin.REQ_OPEN_APP)
}

func CloseApp(stub, deviceID, packName string) error {
	return wrapQueueRequest(stub, deviceID, packName, "close app", bilicoin.REQ_CLOSE_APP)
}

func wrapQueueRequest(stub, deviceID, data, msg string, operation int) error {

	bilicoin.Info("[RMQ] "+msg, logrus.Fields{"stub": stub, "deviceID": deviceID})
	op := MQOrder{
		Id:        deviceID,
		Stub:      stub,
		Data:      data,
		Msg:       msg,
		Operation: operation,
	}

	dat, err := json.Marshal(op)
	if err != nil {
		bilicoin.Fatal("[RMQ] error marshal")
	}

	if err = OrderQueue.Publish(string(dat)); err != nil {
		bilicoin.Fatal("[RMQ] failed to publish")
	}
	return nil
}
