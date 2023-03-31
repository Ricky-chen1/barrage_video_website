package dao

import "barrage_video_website/model"

type forwardDao struct{}

var forwardInstance *forwardDao

func NewForward() *forwardDao {
	if forwardInstance == nil {
		forwardInstance = &forwardDao{}
	}
	return forwardInstance
}

func (*forwardDao) Create(forward *model.Forward) error {
	if err := DB.Create(forward).Error; err != nil {
		return err
	}
	return nil
}
