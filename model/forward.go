package model

import "gorm.io/gorm"

type Forward struct {
	gorm.Model
	Uid uint `gorm:"not null"`
	Vid uint `gorm:"not null"`
}
