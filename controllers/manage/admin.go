package manage

import (
	"barrage_video_website/pkg/errno"
	"barrage_video_website/serializer"
	"barrage_video_website/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminLogin(c *gin.Context) {
	var AdminLogin service.AdminLoginService
	if err := c.ShouldBind(&AdminLogin); err != nil {
		c.JSON(http.StatusBadRequest, serializer.MakeErrResponse(errno.ErrQueryPramsInvalid))
		return
	}
	res := AdminLogin.Login()
	c.JSON(http.StatusOK, res)
}

func AdminAdd(c *gin.Context) {
	var AdminAdd service.AdminAddService
	if err := c.ShouldBind(&AdminAdd); err != nil {
		c.JSON(http.StatusBadRequest, serializer.MakeErrResponse(errno.ErrQueryPramsInvalid))
		return
	}
	res := AdminAdd.Add()
	c.JSON(http.StatusOK, res)
}
