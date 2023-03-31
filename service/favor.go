package service

import (
	"barrage_video_website/dao"
	"barrage_video_website/model"
	"barrage_video_website/pkg/errno"
	"barrage_video_website/serializer"
)

type FavorService struct {
	Vid uint `json:"vid" form:"vid"`
}

// 点赞
func (service *FavorService) Favor(uid uint) serializer.Response {
	if !dao.NewVideo().IsVideoExist(service.Vid) {
		return serializer.MakeErrResponse(errno.ErrVideoNoExist)
	}

	//判断是否已经点赞
	if dao.NewFavor().IsFavor(uid, service.Vid) {
		return serializer.MakeErrResponse(errno.ErrFavorExist)
	}

	NewFavor := model.Favor{
		Uid: uid,
		Vid: service.Vid,
	}
	if err := dao.NewFavor().Create(NewFavor); err != nil {
		return serializer.MakeErrResponse(errno.ErrFavorFail)
	}

	return serializer.Response{
		Status: errno.Success.Code,
		Msg:    errno.Success.Msg,
	}
}
