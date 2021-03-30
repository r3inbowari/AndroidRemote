package service

import "time"

type ResponseMessage struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Ts   int64       `json:"ts"`
}

func SucceedResponse(data interface{}) ResponseMessage {
	return ResponseMessage{
		Code: 200,
		Msg:  "ok",
		Data: data,
		Ts:   time.Now().Unix(),
	}
}

func FailedResponse(msg string) ResponseMessage {
	return ResponseMessage{
		Code: 101,
		Msg:  msg,
		Data: nil,
		Ts:   time.Now().Unix(),
	}
}
