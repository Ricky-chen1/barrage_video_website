package dao

import "barrage_video_website/model"

type adminDao struct{}

var adminInstance *adminDao

// 单例模式创建对象
func NewAdmin() *adminDao {
	if adminInstance == nil {
		adminInstance = &adminDao{}
	}
	return adminInstance
}

func (*adminDao) QueryAdmin(conditions map[string]interface{}) (*model.Admin, error) {
	var admin model.Admin
	if err := DB.Model(&model.Admin{}).Where(conditions).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

func (*adminDao) IsAdminExist(email string) (*model.Admin, bool) {
	var (
		admin model.Admin
		count int64
	)
	DB.Model(&model.Admin{}).Where("email=?", email).First(&admin).Count(&count)
	if count != 0 {
		return &admin, true
	} else {
		return nil, false
	}
}

func (*adminDao) Create(values map[string]interface{}) error {
	if err := DB.Model(&model.Admin{}).Create(values).Error; err != nil {
		return err
	}
	return nil
}
