package controllers

import (
	"barrage_video_website/pkg/errno"
	"barrage_video_website/serializer"
	"barrage_video_website/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 搜索视频
func SearchVideo(c *gin.Context) {
	var SearchVideo service.SearchVideoService
	if err := c.ShouldBindQuery(&SearchVideo); err != nil {
		c.JSON(http.StatusBadRequest, serializer.MakeErrResponse(errno.ErrQueryPramsInvalid))
	}

	res := SearchVideo.SearchVideo()
	c.JSON(http.StatusOK, res)
}
