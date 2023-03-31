package controllers

import (
	"barrage_video_website/pkg/errno"
	"barrage_video_website/pkg/util"
	"barrage_video_website/serializer"
	"barrage_video_website/service"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 上传视频
func VideoUpload(c *gin.Context) {
	var UploadService *service.VideoUploadService
	if err := c.ShouldBind(&UploadService); err != nil {
		c.JSON(http.StatusBadRequest, serializer.MakeErrResponse(errno.ErrQueryPramsInvalid))
		return
	}

	//获取要上传的视频
	video, err := c.FormFile("video")
	if err != nil {
		c.JSON(http.StatusBadRequest, serializer.MakeErrResponse(errno.ErrVideoUpload))
		return
	}

	//检查视频类型
	videoExt := path.Ext(video.Filename)
	if videoExt != ".mp4" {
		c.JSON(http.StatusUnprocessableEntity, serializer.MakeErrResponse(errno.ErrFileTypeInvalid))
		return
	}

	//生成视频要保存的目录
	dir := "./upload/video" + util.GetDay()
	if _, err := os.Stat(dir); err != nil {

		//生成视频保存目录
		if err := os.MkdirAll(dir, 0666); err != nil {
			c.JSON(http.StatusInternalServerError, serializer.MakeErrResponse(errno.ErrSystem))
			return
		}
	}

	//生成视频名称
	videoName := util.RandStringByte(3) + strconv.FormatInt(util.GetUnix(), 10)
	video.Filename = videoName + videoExt

	//上传文件至指定路径
	dst := path.Join(dir, video.Filename)
	if err := c.SaveUploadedFile(video, dst); err != nil {
		c.JSON(http.StatusBadRequest, serializer.MakeErrResponse(errno.ErrSaveFileFail))
		return
	}

	//拿到用户id并调用service层服务
	uid := c.GetUint("uid")
	res := UploadService.Upload(uid, dst)
	c.JSON(http.StatusOK, res)
}

// 获取用户个人视频详情
func ShowVideo(c *gin.Context) {
	var ShowVideo service.VideoShowService
	if err := c.ShouldBindQuery(&ShowVideo); err != nil {
		c.JSON(http.StatusBadRequest, serializer.MakeErrResponse(errno.ErrQueryPramsInvalid))
		return
	}

	uid := c.GetUint("uid")
	res := ShowVideo.Show(uid)
	c.JSON(http.StatusOK, res)
}
