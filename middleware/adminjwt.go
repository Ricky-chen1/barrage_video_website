package middleware

import (
	"barrage_video_website/pkg/errno"
	"barrage_video_website/pkg/util"
	"barrage_video_website/serializer"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func AdminJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, serializer.MakeErrResponse(errno.ErrNoToken))
			c.Abort()
			return
		}
		admin_claims, err := util.ParseAdminToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, serializer.MakeErrResponse(errno.ErrTokenParseFail))
			c.Abort()
			return
		} else if time.Now().Unix() > admin_claims.ExpiresAt {
			c.JSON(http.StatusUnauthorized, serializer.MakeErrResponse(errno.ErrTokenExpired))
			c.Abort()
			return
		}

		c.Next()
	}
}
