package dao

import "barrage_video_website/model"

type replyDao struct{}

var replyInstance *replyDao

// 单例模式创建对象
func NewReply() *replyDao {
	if replyInstance == nil {
		replyInstance = &replyDao{}
	}
	return replyInstance
}

func (*replyDao) Create(reply *model.Reply) error {
	if err := DB.Create(reply).Error; err != nil {
		return err
	}
	return nil
}
