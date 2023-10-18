package models

import "net/http"

var (
	OK = Response{
		Code: http.StatusOK,
		Msg:  "操作成功",
		Data: nil,
	}
	ERR = Response{
		Code: http.StatusInternalServerError,
		Msg:  "操作失败",
		Data: nil,
	}
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func (r Response) WithMsg(message string) Response {
	return Response{
		Code: r.Code,
		Msg:  message,
		Data: r.Data,
	}
}

func (r Response) WithData(data interface{}) Response {
	return Response{
		Code: r.Code,
		Msg:  r.Msg,
		Data: data,
	}
}

func (r Response) WithCode(code int) Response {
	return Response{
		Code: code,
		Msg:  r.Msg,
		Data: r.Data,
	}
}
