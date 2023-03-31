package model

import "gorm.io/gorm"

type Barrage struct {
	gorm.Model
	Uid     uint   `gorm:"not null"`
	Vid     uint   `gorm:"not null"`
	Content string `gorm:"not null"`
	Type    int    `gorm:"default:0"` //-1底部,0滚动,1顶部,
}
