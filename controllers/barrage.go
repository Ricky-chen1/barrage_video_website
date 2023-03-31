package controllers

import (
	"barrage_video_website/pkg/errno"
	"barrage_video_website/serializer"
	"barrage_video_website/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 发送弹幕
func SendBarrage(c *gin.Context) {
	var sendBarrage service.SendBarrageService
	if err := c.ShouldBind(&sendBarrage); err != nil {
		c.JSON(http.StatusBadRequest, serializer.MakeErrResponse(errno.ErrQueryPramsInvalid))
		return
	}

	uid := c.GetUint("uid")
	res := sendBarrage.SendBarrage(uid)
	c.JSON(http.StatusOK, res)
}

// 获取弹幕列表
func GetBarrageList(c *gin.Context) {
	var getBarrageList service.GetBarrageListService
	if err := c.ShouldBindQuery(&getBarrageList); err != nil {
		c.JSON(http.StatusBadRequest, serializer.MakeErrResponse(errno.ErrQueryPramsInvalid))
		return
	}

	res := getBarrageList.GetList()
	c.JSON(http.StatusOK, res)
}
