package service

import (
	"barrage_video_website/cache"
	"barrage_video_website/dao"
	"barrage_video_website/pkg/errno"
	"barrage_video_website/serializer"
)

type SearchVideoService struct {
	Keywords string `json:"keywords" form:"keywords"`
	PageSize int    `json:"page_size" form:"page_size"`
	PageNum  int    `json:"page_num" form:"page_num"`
}

func (service *SearchVideoService) SearchVideo() serializer.Response {

	videos, err := dao.NewSearch().SearchVideo(service.Keywords, service.PageSize, service.PageNum)
	if err != nil {
		return serializer.MakeErrResponse(errno.ErrSearchVideoFail)
	}

	//是否有匹配的结果
	if len(videos) == 0 {
		return serializer.MakeErrResponse(errno.ErrNoSearchVideo)
	}

	//缓存搜索记录
	cache.RedisNewClient.LPush(cache.Ctx, cache.SearchVideoKey, service.Keywords)

	return serializer.Response{
		Status: errno.Success.Code,
		Msg:    errno.Success.Msg,
		Data:   serializer.BuildVideos(videos),
	}
}
