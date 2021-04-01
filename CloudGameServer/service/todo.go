package service

import (
	bilicoin "CloudGameServer/utils"
	"time"
)

type ResponseMessage struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Ts   int64       `json:"ts"`
}

func SucceedResponse(data interface{}, code ...int) ResponseMessage {
	var code1 = bilicoin.Normal
	if len(code) > 0 {
		code1 = code[0]
	}
	return ResponseMessage{
		Code: code1,
		Msg:  "ok",
		Data: data,
		Ts:   time.Now().Unix(),
	}
}

func FailedResponse(msg string, code ...int) ResponseMessage {
	var code1 = bilicoin.InternalServerError
	if len(code) > 0 {
		code1 = code[0]
	}
	return ResponseMessage{
		Code: code1,
		Msg:  msg,
		Data: nil,
		Ts:   time.Now().Unix(),
	}
}
