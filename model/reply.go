package model

import "gorm.io/gorm"

type Reply struct {
	gorm.Model
	Cid     uint `gorm:"default:0"` //回复的评论id
	Content string
	Uid     uint
}
