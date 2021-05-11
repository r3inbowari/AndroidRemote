package bilicoin

import "time"

type ResponseBody struct {
	Code int
	Msg  string
	Data interface{}
	Ts   int64
}

type RequestBody struct {
	Code int
	Msg  string
	Date string
	Ts   int64
}

func ResponseOKWrap(msg string) ResponseBody {

	return ResponseBody{
		Code: 200,
		Msg:  msg,
		Data: "",
		Ts:   time.Now().Unix(),
	}
}

func ResponseWrap(code int, msg string) ResponseBody {

	return ResponseBody{
		Code: code,
		Msg:  msg,
		Data: "",
		Ts:   time.Now().Unix(),
	}
}

func ResponseWrapWithDate(code int, msg string, data interface{}) ResponseBody {
	return ResponseBody{
		Code: code,
		Msg:  msg,
		Data: data,
		Ts:   time.Now().Unix(),
	}
}

var HashValue = "dev"
