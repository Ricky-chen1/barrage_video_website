package model

import (
	"barrage_video_website/cache"
	"strconv"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Uid     uint `gorm:"index"`
	AddTime int64
	Dst     string `gorm:"not null"`
	Title   string `gorm:"index"`
	Review  bool   `gorm:"not null"` //视频审核状态
}

func (video *Video) VideoViewGet() uint64 {
	//获取视频点击量
	count, _ := cache.RedisNewClient.Get(cache.Ctx, cache.VideoViewKey(video.ID)).Uint64()
	return count
}

// 浏览视频
func (video *Video) VideoViewAdd() {
	//增加视频点击量
	cache.RedisNewClient.Incr(cache.Ctx, cache.VideoViewKey(video.ID))
	//增加排行榜点击量
	cache.RedisNewClient.ZIncrBy(cache.Ctx, cache.RankDailyKey, 1, strconv.Itoa(int(video.ID)))
}
