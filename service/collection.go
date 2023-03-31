package service

import (
	"barrage_video_website/dao"
	"barrage_video_website/model"
	"barrage_video_website/pkg/errno"
	"barrage_video_website/serializer"
)

type CreateCollectionService struct {
	Name   string `json:"name" form:"name"`
	IsOpen bool   `json:"is_open" form:"is_open"`
}

// 创建收藏夹
func (service *CreateCollectionService) Create(uid uint) serializer.Response {
	collection := model.Collection{
		Uid:    uid,
		Name:   service.Name,
		IsOpen: service.IsOpen,
	}

	if err := dao.NewCollection().Create(collection); err != nil {
		return serializer.MakeErrResponse(errno.ErrCollectionCreateFail)
	}

	return serializer.Response{
		Status: errno.Success.Code,
		Msg:    errno.Success.Msg,
		Data:   serializer.BuildCollection(collection),
	}
}
