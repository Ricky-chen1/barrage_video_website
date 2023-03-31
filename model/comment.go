package model

import (
	"barrage_video_website/cache"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Vid     uint
	Content string `gorm:"not null"`
	Uid     uint   `gorm:"index"`
}

func (comment *Comment) ReplyAdd() {
	cache.RedisNewClient.Incr(cache.Ctx, cache.CommentReplyKey(comment.ID))
}

func (comment *Comment) GetReplyCount() uint64 {
	count, _ := cache.RedisNewClient.Get(cache.Ctx, cache.CommentReplyKey(comment.ID)).Uint64()
	return count
}
