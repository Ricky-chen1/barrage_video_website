package model

import "gorm.io/gorm"

//收藏
type Collect struct {
	gorm.Model
	CollectionName string `gorm:"not null"`
	Vid            uint   `gorm:"not null"`
	Uid            uint   `gorm:"not null"`
}
