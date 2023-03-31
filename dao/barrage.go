package dao

import "barrage_video_website/model"

type barrageDao struct{}

var barrageInstance *barrageDao

func NewBarrage() *barrageDao {
	if barrageInstance == nil {
		barrageInstance = &barrageDao{}
	}
	return barrageInstance
}

func (*barrageDao) Create(barrage model.Barrage) error {
	if err := DB.Create(&barrage).Error; err != nil {
		return err
	}
	return nil
}

func (*barrageDao) GetList(conditions map[string]interface{}, pageSize int, pageNum int) ([]model.Barrage, error) {
	var barrages []model.Barrage
	err := DB.Model(&model.Barrage{}).Where(conditions).
		Scopes(Paginate(pageSize, pageNum)).Find(&barrages).Error

	if err != nil {
		return nil, err
	}
	return barrages, nil
}
