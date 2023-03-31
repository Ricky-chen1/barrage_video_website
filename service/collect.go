package service

import (
	"barrage_video_website/dao"
	"barrage_video_website/model"
	"barrage_video_website/pkg/errno"
	"barrage_video_website/serializer"
)

type CollectService struct {
	Vid            uint   `json:"vid" form:"vid"`
	CollectionName string `json:"collection_name" form:"collection_name"`
}

// 收藏
func (service *CollectService) Collect(uid uint) serializer.Response {
	if !dao.NewVideo().IsVideoExist(service.Vid) {
		return serializer.MakeErrResponse(errno.ErrVideoNoExist)
	}

	if dao.NewCollect().IsCollect(uid, service.Vid) {
		return serializer.MakeErrResponse(errno.ErrCollectExist)
	}

	//未传值默认选择默认收藏夹
	if len(service.CollectionName) == 0 {
		service.CollectionName = "默认收藏夹"
	}

	//用户是否有该收藏夹
	if !dao.NewCollection().IsCollectionBelongToUser(uid, service.CollectionName) {
		return serializer.MakeErrResponse(errno.ErrCollectionNoExist)
	}

	NewCollect := model.Collect{
		Uid:            uid,
		Vid:            service.Vid,
		CollectionName: service.CollectionName,
	}

	//收藏
	if err := dao.NewCollect().Create(NewCollect); err != nil {
		return serializer.MakeErrResponse(errno.ErrCollectFail)
	}

	return serializer.Response{
		Status: errno.Success.Code,
		Msg:    errno.Success.Msg,
	}
}
