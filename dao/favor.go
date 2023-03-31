package dao

import "barrage_video_website/model"

type favorDao struct{}

var favaoInstance *favorDao

func NewFavor() *favorDao {
	if favaoInstance == nil {
		favaoInstance = &favorDao{}
	}
	return favaoInstance
}

//是否点赞
func (*favorDao) IsFavor(uid uint, vid uint) bool {
	var (
		favor model.Favor
		count int64
	)
	DB.Model(&model.Favor{}).Where("uid=? AND vid=?", uid, vid).First(&favor).Count(&count)
	return count != 0
}

func (*favorDao) Create(favor model.Favor) error {
	if err := DB.Create(&favor).Error; err != nil {
		return err
	}
	return nil
}
