package serializer

import (
	"barrage_video_website/model"
)

type Video struct {
	Id     uint   `json:"id"`
	Uid    uint   `json:"uid"`
	Title  string `json:"title"`
	Dst    string `json:"dst"`
	View   uint   `json:"view"`
	Review bool   `json:"review"` //审核状态
}

// 序列化视频信息
func BuildVideo(video model.Video) Video {
	return Video{
		Id:     video.ID,
		Uid:    video.Uid,
		Title:  video.Title,
		Dst:    video.Dst,
		View:   uint(video.VideoViewGet()),
		Review: video.Review,
	}
}

// 序列化视频列表信息
func BuildVideos(items []model.Video) (videos []Video) {
	for _, item := range items {
		video := BuildVideo(item)
		videos = append(videos, video)
	}
	return videos
}
