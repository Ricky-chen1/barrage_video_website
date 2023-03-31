package controllers

import (
	"barrage_video_website/pkg/errno"
	"barrage_video_website/serializer"
	"barrage_video_website/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 创建收藏夹
func CreateCollection(c *gin.Context) {
	var createCollection service.CreateCollectionService
	if err := c.ShouldBind(&createCollection); err != nil {
		c.JSON(http.StatusBadRequest, serializer.MakeErrResponse(errno.ErrQueryPramsInvalid))
		return
	}

	uid := c.GetUint("uid")
	res := createCollection.Create(uid)
	c.JSON(http.StatusOK, res)
}
