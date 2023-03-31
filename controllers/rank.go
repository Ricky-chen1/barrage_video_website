package controllers

import (
	"barrage_video_website/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DailyRank(c *gin.Context) {
	var dailyRank service.DaliyRankService
	res := dailyRank.Get()
	c.JSON(http.StatusOK, res)
}
