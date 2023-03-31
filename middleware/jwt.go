package middleware

import (
	"barrage_video_website/pkg/errno"
	"barrage_video_website/pkg/util"
	"barrage_video_website/serializer"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// JWT鉴权中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, serializer.MakeErrResponse(errno.ErrNoToken))
			c.Abort()
			return
		}
		claims, err := util.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, serializer.MakeErrResponse(errno.ErrTokenParseFail))
			c.Abort()
			return
		} else if time.Now().Unix() > claims.ExpiresAt || util.CheckIsTokenRefresh(claims) {
			c.JSON(http.StatusUnauthorized, serializer.MakeErrResponse(errno.ErrTokenExpired))
			c.Abort()
			return
		}

		//将用户id放入上下文中,方便后续业务
		c.Set("uid", claims.Id)
		c.Next()
	}
}
