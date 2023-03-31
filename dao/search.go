package dao

import "barrage_video_website/model"

type searchDao struct{}

var searchInstance *searchDao

// 单例模式创建对象
func NewSearch() *searchDao {
	if searchInstance == nil {
		searchInstance = &searchDao{}
	}
	return searchInstance
}

//多条件搜索视频
func (*searchDao) SearchVideo(keywords string, pageSize int, pageNum int) ([]model.Video, error) {
	var videos []model.Video
	err := DB.Model(&model.Video{}).Scopes(SearchApprovedVideo()).
		Where(DB.Where("title LIKE ?", "%"+keywords+"%").
			Or("add_time LIKE ?", "%"+keywords+"%").
			Or("uid LIKE ?", "%"+keywords+"%").
			Or("dst LIKE ?", "%"+keywords+"%")).
		Find(&videos).Scopes(Paginate(pageSize, pageNum)).Error
	if err != nil {
		return nil, err
	}

	return videos, nil
}
