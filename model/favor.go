package model

import "gorm.io/gorm"

//点赞实体
type Favor struct {
	gorm.Model
	Uid uint `gorm:"not null"` //点赞的用户
	Vid uint `gorm:"not null"` //点赞的视频
}
