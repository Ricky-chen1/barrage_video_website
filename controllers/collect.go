package controllers

import (
	"barrage_video_website/pkg/errno"
	"barrage_video_website/serializer"
	"barrage_video_website/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Collect(c *gin.Context) {
	var collectService service.CollectService
	if err := c.ShouldBind(&collectService); err != nil {
		c.JSON(http.StatusBadRequest, serializer.MakeErrResponse(errno.ErrQueryPramsInvalid))
		return
	}

	uid := c.GetUint("uid")
	res := collectService.Collect(uid)
	c.JSON(http.StatusOK, res)
}
