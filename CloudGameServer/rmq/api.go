package rmq

import (
	bilicoin "CloudGameServer/utils"
	"encoding/json"
	"github.com/sirupsen/logrus"
)

func OpenSender(stub, deviceID string) error {

	bilicoin.Info("[WS] open sender", logrus.Fields{"stub": stub, "deviceID": deviceID})
	op := MQOrder{
		Id:        deviceID,
		Stub:      stub,
		Data:      "",
		Msg:       "order",
		Operation: bilicoin.REQ_START_SENDER,
	}

	dat, err := json.Marshal(op)
	if err != nil {
		bilicoin.Fatal("error marshal")
	}

	if err = OrderQueue.Publish(string(dat)); err != nil {
		bilicoin.Fatal("failed to publish")
	}

	return nil
}

func CloseSender(stub, deviceID string) error {

	op := MQOrder{
		Id:        deviceID,
		Stub:      stub,
		Data:      "",
		Msg:       "order",
		Operation: bilicoin.REQ_PAUSE_SENDER,
	}

	dat, err := json.Marshal(op)
	if err != nil {
		bilicoin.Fatal("error marshal")
	}
	if err = OrderQueue.Publish(string(dat)); err != nil {
		bilicoin.Fatal("failed to publish")
	}
	return nil
}

func OpenApp(stub, deviceID, appName string) error {

	op := MQOrder{
		Id:        deviceID,
		Stub:      stub,
		Data:      appName,
		Msg:       "order",
		Operation: bilicoin.REQ_OPEN_APP,
	}

	dat, err := json.Marshal(op)
	if err != nil {
		bilicoin.Fatal("error marshal")
	}

	if err = OrderQueue.Publish(string(dat)); err != nil {
		bilicoin.Fatal("failed to publish")
	}

	return nil
}

func CloseApp(stub, deviceID, packName string) error {

	op := MQOrder{
		Id:   deviceID,
		Stub: stub,
		//Data:      "com.shandagames.falloutshelterUc.uc",
		//"com.xd.ro.roapk"
		Data:      packName,
		Msg:       "order",
		Operation: bilicoin.REQ_CLOSE_APP,
	}

	dat, err := json.Marshal(op)
	if err != nil {
		bilicoin.Fatal("error marshal")
	}

	if err = OrderQueue.Publish(string(dat)); err != nil {
		bilicoin.Fatal("failed to publish")
	}
	return nil
}
