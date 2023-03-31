package service

import (
	"barrage_video_website/cache"
	"barrage_video_website/dao"
	"barrage_video_website/model"
	"barrage_video_website/pkg/errno"
	"barrage_video_website/serializer"
)

type DaliyRankService struct{}

func (service *DaliyRankService) Get() serializer.Response {
	var (
		videos []model.Video
		err    error
	)

	//查询点击数排名前十的视频
	vids := cache.RedisNewClient.ZRevRange(cache.Ctx, cache.RankDailyKey, 0, 9).Val()
	if len(vids) >= 1 {
		videos, err = dao.NewVideo().Order(vids)
		if err != nil {
			return serializer.MakeErrResponse(errno.ErrRankVideosFail)
		}
	}

	return serializer.Response{
		Status: errno.Success.Code,
		Msg:    errno.Success.Msg,
		Data:   serializer.BuildVideos(videos),
	}
}
