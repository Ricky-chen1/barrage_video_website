package dao

import (
	"barrage_video_website/model"
)

type commentDao struct{}

var commentInstance *commentDao

// 单例模式创建对象
func NewComment() *commentDao {
	if commentInstance == nil {
		commentInstance = &commentDao{}
	}
	return commentInstance
}

// 评论是否存在
func (*commentDao) IsCommentExist(cid uint) bool {
	var (
		comment model.Comment
		count   int64
	)

	DB.Model(&model.Comment{}).First(&comment, cid).Count(&count)
	return count != 0
}

// 创建评论表记录
func (*commentDao) Create(comment *model.Comment) error {
	if err := DB.Model(&model.Comment{}).Create(comment).Error; err != nil {
		return err
	}
	return nil
}

// 获取一条评论
func (*commentDao) Get(cid uint) (*model.Comment, error) {
	var comment model.Comment
	if err := DB.Model(&model.Comment{}).Where("id=?", cid).Find(&comment).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

// 获取评论列表
func (*commentDao) GetList(conditions map[string]interface{}, pageSize int, pageNum int) ([]model.Comment, error) {
	var comments []model.Comment
	err := DB.Model(&model.Comment{}).Where(conditions).
		Scopes(Paginate(pageSize, pageNum)).Find(&comments).Error

	if err != nil {
		return nil, err
	}
	return comments, nil
}

// 删除评论
func (*commentDao) Delete(cid uint) error {
	var comment model.Comment
	if err := DB.Where("id=?", cid).Delete(&comment).Error; err != nil {
		return err
	}
	return nil
}
