package model

import (
	"barrage_video_website/cache"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username       string `gorm:"unique;not null"`
	Avator         string
	PasswordDigest string
	Email          string `gorm:"not null"`
}

// 将用户添加至黑名单
func (user *User) BlacklistAdd() {
	cache.RedisNewClient.SAdd(cache.Ctx, cache.UserBlacklistKey, cache.UserKey(user.ID))
}

// 用户是否被拉黑
func (user *User) IsBlackList() bool {
	return cache.RedisNewClient.SIsMember(cache.Ctx, cache.UserBlacklistKey, cache.UserKey(user.ID)).Val()
}

func (user *User) BlacklistDelete() {
	cache.RedisNewClient.SRem(cache.Ctx, cache.UserBlacklistKey, cache.UserKey(user.ID))
}
