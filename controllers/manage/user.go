package manage

import (
	"barrage_video_website/pkg/errno"
	"barrage_video_website/serializer"
	"barrage_video_website/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 管理员拉黑用户
func UserBlacklist(c *gin.Context) {
	var UserBlacklist service.UserBlacklistService
	if err := c.ShouldBind(&UserBlacklist); err != nil {
		c.JSON(http.StatusBadRequest, serializer.MakeErrResponse(errno.ErrQueryPramsInvalid))
		return
	}

	res := UserBlacklist.Blacklist()
	c.JSON(http.StatusOK, res)
}

// 将用户从黑名单中删除
func UserWhitelist(c *gin.Context) {
	var UserWhitelist service.UserWhitelistService
	if err := c.ShouldBind(&UserWhitelist); err != nil {
		c.JSON(http.StatusBadRequest, serializer.MakeErrResponse(errno.ErrQueryPramsInvalid))
		return
	}

	res := UserWhitelist.Whitelist()
	c.JSON(http.StatusOK, res)
}
