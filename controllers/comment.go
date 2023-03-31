package controllers

import (
	"barrage_video_website/pkg/errno"
	"barrage_video_website/serializer"
	"barrage_video_website/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 评论
func CommentAdd(c *gin.Context) {
	var commentAdd service.CommentAddService
	if err := c.ShouldBind(&commentAdd); err != nil {
		c.JSON(http.StatusBadRequest, serializer.MakeErrResponse(errno.ErrQueryPramsInvalid))
		return
	}

	uid := c.GetUint("uid")
	res := commentAdd.Add(uid)
	c.JSON(http.StatusOK, res)
}

// 获取评论列表
func CommentListGet(c *gin.Context) {
	var commentListGet service.CommentListGetService
	if err := c.ShouldBindQuery(&commentListGet); err != nil {
		c.JSON(http.StatusBadRequest, serializer.MakeErrResponse(errno.ErrQueryPramsInvalid))
		return
	}

	res := commentListGet.GetList()
	c.JSON(http.StatusOK, res)
}
