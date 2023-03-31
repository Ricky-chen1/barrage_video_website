package dao

import "barrage_video_website/model"

type collectDao struct{}

var collectInstance *collectDao

func NewCollect() *collectDao {
	if collectInstance == nil {
		collectInstance = &collectDao{}
	}
	return collectInstance
}

func (*collectDao) IsCollect(uid uint, vid uint) bool {
	var (
		collect model.Collect
		count   int64
	)
	DB.Model(&model.Collect{}).Where("uid=? AND vid=?", uid, vid).First(&collect).Count(&count)
	return count != 0
}

func (*collectDao) Create(collect model.Collect) error {
	if err := DB.Create(&collect).Error; err != nil {
		return err
	}
	return nil
}
