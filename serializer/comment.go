package serializer

import (
	"barrage_video_website/model"
)

type Comment struct {
	Id         uint   `json:"id"`
	Content    string `json:"content"`
	ReplyCount uint   `json:"reply_count"`
	Uid        uint   `json:"uid"`
	Vid        uint   `json:"vid"`
}

func BuildComment(comment model.Comment) Comment {
	return Comment{
		Id:         comment.ID,
		Content:    comment.Content,
		Vid:        comment.Vid,
		Uid:        comment.Uid,
		ReplyCount: uint(comment.GetReplyCount()),
	}
}

func BuildComments(items []model.Comment) (comments []Comment) {
	for _, item := range items {
		comment := BuildComment(item)
		comments = append(comments, comment)
	}
	return comments
}
