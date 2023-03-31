package service

import (
	"barrage_video_website/dao"
	"barrage_video_website/model"
	"barrage_video_website/pkg/errno"
	"barrage_video_website/serializer"
)

type BarrgeCommonService struct {
	Vid uint `json:"vid" form:"vid"`
}

type GetBarrageListService struct {
	BarrgeCommonService
	PageSize int `json:"page_size" form:"page_size"`
	PageNum  int `json:"page_num" form:"page_num"`
}

type SendBarrageService struct {
	Content string `json:"content" form:"content"`
	BarrgeCommonService
}

// 发送弹幕
func (service *SendBarrageService) SendBarrage(uid uint) serializer.Response {
	newBarrage := model.Barrage{
		Content: service.Content,
		Vid:     service.Vid,
		Uid:     uid,
	}
	if err := dao.NewBarrage().Create(newBarrage); err != nil {
		return serializer.MakeErrResponse(errno.ErrSendBarrageFail)
	}

	return serializer.Response{
		Status: errno.Success.Code,
		Msg:    errno.Success.Msg,
	}
}

// 获取弹幕列表
func (service *GetBarrageListService) GetList() serializer.Response {
	if !dao.NewVideo().IsVideoExist(service.Vid) {
		return serializer.MakeErrResponse(errno.ErrVideoNoExist)
	}

	barrages, err := dao.NewBarrage().GetList(map[string]interface{}{
		"vid": service.Vid,
	}, service.PageSize, service.PageNum)
	if err != nil {
		return serializer.MakeErrResponse(errno.ErrQueryBarrageListFail)
	}

	return serializer.Response{
		Status: errno.Success.Code,
		Msg:    errno.Success.Msg,
		Data:   serializer.BuildBarrages(barrages),
	}
}
