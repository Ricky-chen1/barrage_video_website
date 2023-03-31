package dao

import (
	"barrage_video_website/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB
var Newlogger logger.Interface

// 数据库连接
func MysqlInit(dsn string) {
	//切记不能为DB,err:=...
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   Newlogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //启用单数表名
		},
	})
	if err != nil {
		panic(err)
	}
	DB = db
	if err := DB.AutoMigrate(&model.User{}, &model.Video{},
		&model.Admin{}, &model.Comment{}, &model.Reply{},
		&model.Favor{}, &model.Collection{}, &model.Collect{},
		&model.Barrage{}, &model.Forward{}); err != nil {
		panic(err)
	}
}

// 分页器
func Paginate(pageSize int, pageNum int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageNum == 0 {
			pageNum = 1
		}
		if pageSize <= 0 {
			pageSize = 10
		}
		offset := (pageNum - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

// 默认获取通过审核的视频
func SearchApprovedVideo() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("review=?", true)
	}
}
