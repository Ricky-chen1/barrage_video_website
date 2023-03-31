package model

import "gorm.io/gorm"

//收藏夹
type Collection struct {
	gorm.Model
	Uid    uint   `gorm:"not null"`
	Name   string `gorm:"index;not null"`
	IsOpen bool   `gorm:"default:true"` //是否公开
}
