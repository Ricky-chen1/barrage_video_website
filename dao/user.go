package dao

import (
	"barrage_video_website/model"
)

type userDao struct{}

var userInstance *userDao

// 单例模式创建对象
func NewUser() *userDao {
	if userInstance == nil {
		userInstance = &userDao{}
	}
	return userInstance
}

// 查询用户是否存在
func (*userDao) IsUserExist(uid interface{}) bool {
	var (
		user  model.User
		count int64
	)
	DB.Model(&model.User{}).Where("id=?", uid).First(&user).Count(&count)
	return count != 0
}

// 数据库查询用户信息
func (*userDao) QueryUser(conditions map[string]interface{}) (*model.User, error) {
	var user model.User
	err := DB.Model(&model.User{}).Where(conditions).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// 数据库创建操作
func (*userDao) Create(user model.User) (*model.User, error) {
	if err := DB.Model(&model.User{}).Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// 数据库更改操作
func (*userDao) Update(uid uint, values map[string]interface{}) error {
	if err := DB.Model(&model.User{}).Where("id=?", uid).Updates(values).Error; err != nil {
		return err
	}
	return nil
}
