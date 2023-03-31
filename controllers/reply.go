package controllers

import (
	"barrage_video_website/pkg/errno"
	"barrage_video_website/serializer"
	"barrage_video_website/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Reply(c *gin.Context) {
	var replyAddService service.ReplyAddService
	if err := c.ShouldBind(&replyAddService); err != nil {
		c.JSON(http.StatusBadRequest, serializer.MakeErrResponse(errno.ErrQueryPramsInvalid))
		return
	}

	uid := c.GetUint("uid")
	res := replyAddService.Add(uid)
	c.JSON(http.StatusOK, res)
}
