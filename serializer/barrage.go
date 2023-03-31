package serializer

import (
	"barrage_video_website/model"
	"barrage_video_website/pkg/util"
)

type Barrage struct {
	Id      uint   `json:"id"`
	Content string `json:"content"`
	Type    int    `json:"type"`
	Uid     uint   `json:"uid"`
	Vid     uint   `json:"vid"`
	Time    int64  `json:"time"`
}

func BuildBarrage(comment model.Barrage) Barrage {
	return Barrage{
		Id:      comment.ID,
		Content: comment.Content,
		Vid:     comment.Vid,
		Uid:     comment.Uid,
		Type:    comment.Type,
		Time:    util.GetUnix(),
	}
}

func BuildBarrages(items []model.Barrage) (barrages []Barrage) {
	for _, item := range items {
		barrage := BuildBarrage(item)
		barrages = append(barrages, barrage)
	}
	return barrages
}
