package controllers

import (
	"barrage_video_website/pkg/errno"
	"barrage_video_website/serializer"
	"barrage_video_website/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Favor(c *gin.Context) {
	var favor service.FavorService
	if err := c.ShouldBind(&favor); err != nil {
		c.JSON(http.StatusBadRequest, serializer.MakeErrResponse(errno.ErrQueryPramsInvalid))
		return
	}

	uid := c.GetUint("uid")
	res := favor.Favor(uid)
	c.JSON(http.StatusOK, res)
}
