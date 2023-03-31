package serializer

import (
	"barrage_video_website/pkg/errno"
)

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Msg    string      `json:"msg"`
}

// 错误序列化函数
func MakeErrResponse(e *errno.Errno) Response {
	return Response{
		Status: e.Code,
		Msg:    e.Message,
	}
}
