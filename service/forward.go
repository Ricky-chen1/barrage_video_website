package service

import (
	"barrage_video_website/dao"
	"barrage_video_website/model"
	"barrage_video_website/pkg/errno"
	"barrage_video_website/serializer"
)

type VideoForwardService struct {
	Vid uint `json:"vid" form:"vid"`
}

// 转发视频
func (service *VideoForwardService) Forward(uid uint) serializer.Response {
	if !dao.NewVideo().IsVideoExist(service.Vid) {
		return serializer.MakeErrResponse(errno.ErrVideoNoExist)
	}

	NewForward := &model.Forward{
		Vid: service.Vid,
		Uid: uid,
	}

	if err := dao.NewForward().Create(NewForward); err != nil {
		return serializer.MakeErrResponse(errno.ErrForwardFail)
	}

	return serializer.Response{
		Status: errno.Success.Code,
		Msg:    errno.Success.Msg,
	}
}
