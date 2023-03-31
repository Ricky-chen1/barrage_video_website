package dao

import "barrage_video_website/model"

type collectionDao struct{}

var collectionInstance *collectionDao

func NewCollection() *collectionDao {
	if collectionInstance == nil {
		collectionInstance = &collectionDao{}
	}
	return collectionInstance
}

func (*collectionDao) IsCollectionExist(id uint) bool {
	var (
		collection model.Collection
		count      int64
	)

	DB.Model(&model.Collection{}).First(&collection, id).Count(&count)
	return count != 0
}

func (*collectionDao) Create(colletion model.Collection) error {
	if err := DB.Create(&colletion).Error; err != nil {
		return err
	}
	return nil
}

func (*collectionDao) IsCollectionBelongToUser(uid uint, name string) bool {
	var count int64
	DB.Model(&model.Collection{}).Where("uid = ? AND name = ?", uid, name).Count(&count)
	return count != 0
}
