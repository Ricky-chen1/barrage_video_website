package dao

import (
	"barrage_video_website/model"
	"fmt"
	"strings"
)

type videoDao struct{}

var videoInstance *videoDao

// 单例模式创建对象
func NewVideo() *videoDao {
	if videoInstance == nil {
		videoInstance = &videoDao{}
	}
	return videoInstance
}

func (*videoDao) IsVideoExist(vid uint) bool {
	var (
		video model.Video
		count int64
	)
	DB.Model(&model.Video{}).Where("id=?", vid).First(&video).Count(&count)
	return count != 0
}

func (*videoDao) Find(vid uint, uid uint) (*model.Video, error) {
	var video model.Video
	if err := DB.Model(&model.Video{}).Scopes(SearchApprovedVideo()).
		Where("id=? AND uid=?", vid, uid).First(&video).Error; err != nil {
		return nil, err
	}
	return &video, nil
}

func (*videoDao) Create(values map[string]interface{}) error {
	if err := DB.Model(&model.Video{}).Create(values).Error; err != nil {
		return err
	}
	return nil
}

func (*videoDao) Update(vid uint, values map[string]interface{}) error {
	if err := DB.Model(&model.Video{}).Where("id=?", vid).Updates(values).Error; err != nil {
		return err
	}
	return nil
}

func (*videoDao) Order(vids []string) ([]model.Video, error) {
	var videos []model.Video
	order := fmt.Sprintf("FIELD(id,%s)", strings.Join(vids, ","))
	err := DB.Model(model.Video{}).Where("id in (?)", vids).Order(order).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}
