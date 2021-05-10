package rmq

import (
	bilicoin "CloudGameServer/utils"
	"encoding/json"
	"time"
)

func OpenSender(stub, appName string) error {

	//// 检查stub
	//CheckStub()
	//
	//// 发送命令
	//Order()

	op := MQOrder{
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

	if err = OrderQueue.Publish(string(dat)); err != nil {
		bilicoin.Fatal("failed to publish")
	}

	return nil
}

func CloseSender(stub, appName string) error {

	//// 检查stub
	//CheckStub()
	//
	//// 发送命令
	//Order()

	op := MQOrder{
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
	if err = OrderQueue.Publish(string(dat)); err != nil {
		bilicoin.Fatal("failed to publish")
	}
	return nil
}

func OpenApp(stub, appName string) error {

	op := MQOrder{
		Id:        "deab9dbaaa74541d",
		Stub:      bilicoin.CreateMD5(time.Now().String()),
		Data:      "仙境传说RO",
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

func CloseApp(stub, appName string) error {

	op := MQOrder{
		Id:        "deab9dbaaa74541d",
		Stub:      bilicoin.CreateMD5(time.Now().String()),
		//Data:      "com.shandagames.falloutshelterUc.uc",
		Data: "com.xd.ro.roapk",
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
