package service

import (
	"barrage_video_website/dao"
	"barrage_video_website/model"
	"barrage_video_website/pkg/errno"
	"barrage_video_website/serializer"
)

type CommentCommonService struct {
	Vid uint `json:"vid" form:"vid"`
}

type CommentAddService struct {
	Content string `json:"content" form:"content"`
	CommentCommonService
}

type CommentListGetService struct {
	PageSize int `json:"page_size" form:"page_size"`
	PageNum  int `json:"page_num" form:"page_num"`
	CommentCommonService
}

type CommentDeleteService struct{}

// 评论
func (service *CommentAddService) Add(uid uint) serializer.Response {
	//查询要评论的视频是否存在
	if !dao.NewVideo().IsVideoExist(service.Vid) {
		return serializer.MakeErrResponse(errno.ErrVideoNoExist)
	}

	comment := model.Comment{
		Uid:     uid,
		Content: service.Content,
		Vid:     service.Vid,
	}
	//增加评论
	if err := dao.NewComment().Create(&comment); err != nil {
		return serializer.MakeErrResponse(errno.ErrCommentAddFail)
	}

	return serializer.Response{
		Status: errno.Success.Code,
		Msg:    errno.Success.Msg,
		Data:   serializer.BuildComment(comment),
	}
}

// 获取评论列表
func (service *CommentListGetService) GetList() serializer.Response {
	if !dao.NewVideo().IsVideoExist(service.Vid) {
		return serializer.MakeErrResponse(errno.ErrVideoNoExist)
	}

	comments, err := dao.NewComment().GetList(map[string]interface{}{
		"vid": service.Vid,
	}, service.PageSize, service.PageNum)
	if err != nil {
		return serializer.MakeErrResponse(errno.ErrQueryCommentListFail)
	}

	return serializer.Response{
		Status: errno.Success.Code,
		Msg:    errno.Success.Msg,
		Data:   serializer.BuildComments(comments),
	}
}

// 管理员删除视频评论
func (service *CommentDeleteService) Delete(cid uint) serializer.Response {
	if !dao.NewComment().IsCommentExist(cid) {
		return serializer.MakeErrResponse(errno.ErrCommentNoExist)
	}

	if err := dao.NewComment().Delete(cid); err != nil {
		return serializer.MakeErrResponse(errno.ErrCommentDelFail)
	}
	return serializer.Response{
		Status: errno.Success.Code,
		Msg:    errno.Success.Msg,
	}
}
