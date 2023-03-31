package service

import (
	"barrage_video_website/dao"
	"barrage_video_website/pkg/errno"
	"barrage_video_website/pkg/util"
	"barrage_video_website/serializer"
)

type VideoReviewService struct {
	Vid    uint `json:"vid" form:"vid"`
	Status bool `json:"status" form:"status"`
}

type VideoUploadService struct {
	Title string `json:"title" form:"title"`
}

type VideoShowService struct {
	Vid uint `json:"vid" form:"vid"`
}

// 上传视频
func (service *VideoUploadService) Upload(uid uint, dst string) serializer.Response {
	if err := dao.NewVideo().Create(map[string]interface{}{
		"title":    service.Title,
		"uid":      uid,
		"dst":      dst, //本地保存路径
		"add_time": util.GetUnix(),
		"review":   false,
	}); err != nil {
		return serializer.MakeErrResponse(errno.ErrCreateVideoRecordFail)
	}

	return serializer.Response{
		Status: errno.Success.Code,
		Msg:    errno.Success.Msg,
	}
}

// 获取用户个人视频详情
func (service *VideoShowService) Show(uid uint) serializer.Response {
	if !dao.NewVideo().IsVideoExist(service.Vid) {
		return serializer.MakeErrResponse(errno.ErrVideoNoExist)
	}

	video, err := dao.NewVideo().Find(service.Vid, uid)
	if err != nil {
		return serializer.MakeErrResponse(errno.ErrQueryVideosFail)
	}

	//增加视频点击数
	video.VideoViewAdd()

	return serializer.Response{
		Status: errno.Success.Code,
		Msg:    errno.Success.Msg,
		Data:   serializer.BuildVideo(*video),
	}
}

// 视频审核
func (service *VideoReviewService) Review() serializer.Response {
	if !dao.NewVideo().IsVideoExist(service.Vid) {
		return serializer.MakeErrResponse(errno.ErrVideoNoExist)
	}

	//更新视频审核状态
	if err := dao.NewVideo().Update(service.Vid, map[string]interface{}{
		"review": service.Status,
	}); err != nil {
		return serializer.MakeErrResponse(errno.ErrReviewVideosFail)
	}

	return serializer.Response{
		Status: errno.Success.Code,
		Msg:    errno.Success.Msg,
	}
}
