package model

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	AdminName      string `gorm:"default:'管理员'"`
	PasswordDigest string `gorm:"not null"`
	Email          string `gorm:"unique;index;not null"`
	Authority      int    `gorm:"not null" `
}
