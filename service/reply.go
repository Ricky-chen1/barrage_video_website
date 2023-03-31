package service

import (
	"barrage_video_website/dao"
	"barrage_video_website/model"
	"barrage_video_website/pkg/errno"
	"barrage_video_website/serializer"
)

type ReplyAddService struct {
	Content string `json:"content" form:"content"`
	Cid     uint   `json:"cid" form:"cid"`
}

// 评论回复
func (service *ReplyAddService) Add(uid uint) serializer.Response {
	if !dao.NewComment().IsCommentExist(service.Cid) {
		return serializer.MakeErrResponse(errno.ErrCommentNoExist)
	}

	reply := model.Reply{
		Cid:     service.Cid,
		Content: service.Content,
		Uid:     uid,
	}
	if err := dao.NewReply().Create(&reply); err != nil {
		return serializer.MakeErrResponse(errno.ErrCommentReplyFail)
	}

	//增加评论回复数
	comment, err := dao.NewComment().Get(service.Cid)
	if err != nil {
		return serializer.MakeErrResponse(errno.ErrQueryCommentFail)
	}
	comment.ReplyAdd()

	return serializer.Response{
		Status: errno.Success.Code,
		Msg:    errno.Success.Msg,
	}
}
