package controllers

import (
	"barrage_video_website/pkg/errno"
	"barrage_video_website/pkg/util"
	"barrage_video_website/serializer"
	"barrage_video_website/service"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var RegisterService service.UserRegisterService
	if err := c.ShouldBind(&RegisterService); err != nil {
		c.JSON(http.StatusBadRequest, serializer.MakeErrResponse(errno.ErrQueryPramsInvalid))
		return
	}

	res := RegisterService.Register()
	c.JSON(http.StatusOK, res)
}

func UserLogin(c *gin.Context) {
	var LoginService service.UserLoginService
	if err := c.ShouldBind(&LoginService); err != nil {
		c.JSON(http.StatusBadRequest, serializer.MakeErrResponse(errno.ErrQueryPramsInvalid))
		return
	}

	res := LoginService.Login()
	c.JSON(http.StatusOK, res)
}

func UserShow(c *gin.Context) {
	var ShowService service.UserShowService
	if err := c.ShouldBind(&ShowService); err != nil {
		c.JSON(http.StatusBadRequest, serializer.MakeErrResponse(errno.ErrQueryPramsInvalid))
		return
	}

	//从上下文中拿出鉴权用户id
	uid := c.GetUint("uid")
	res := ShowService.Show(uid)
	c.JSON(http.StatusOK, res)
}

func UserUpdateName(c *gin.Context) {
	var UpdateNameService service.UserUpdateNameService
	if err := c.ShouldBind(&UpdateNameService); err != nil {
		c.JSON(http.StatusBadRequest, serializer.MakeErrResponse(errno.ErrQueryPramsInvalid))
		return
	}

	//从上下文中拿出鉴权用户id
	uid := c.GetUint("uid")
	res := UpdateNameService.UpdateName(uid)
	c.JSON(http.StatusOK, res)
}

func UserUpdateEmail(c *gin.Context) {
	var UpdateEmailService service.UserUpdateEmailService
	if err := c.ShouldBind(&UpdateEmailService); err != nil {
		c.JSON(http.StatusBadRequest, serializer.MakeErrResponse(errno.ErrQueryPramsInvalid))
		return
	}

	//从上下文中拿出鉴权用户id
	uid := c.GetUint("uid")
	res := UpdateEmailService.UpdateEmail(uid)
	c.JSON(http.StatusOK, res)
}

func UserModifyPassword(c *gin.Context) {
	var ModifyPasswordService service.UserModifyPasswordService
	if err := c.ShouldBind(&ModifyPasswordService); err != nil {
		c.JSON(http.StatusBadRequest, serializer.MakeErrResponse(errno.ErrQueryPramsInvalid))
		return
	}

	//从上下文中拿出鉴权用户id
	uid := c.GetUint("uid")
	res := ModifyPasswordService.Modify(uid)
	c.JSON(http.StatusOK, res)
}

func UserAvatorUpload(c *gin.Context) {
	var UploadAvatorService service.UserAvatorUploadService
	if err := c.ShouldBind(&UploadAvatorService); err != nil {
		c.JSON(http.StatusBadRequest, serializer.MakeErrResponse(errno.ErrQueryPramsInvalid))
		return
	}

	//获取要上传的头像图片
	avator, err := c.FormFile("avator")
	if err != nil {
		c.JSON(http.StatusBadRequest, serializer.MakeErrResponse(errno.ErrGetFileFail))
		return
	}

	//检查文件格式
	avatorExt := path.Ext(avator.Filename)
	if avatorExt != ".jpg" && avatorExt != ".png" && avatorExt != ".jpeg" {
		c.JSON(http.StatusUnprocessableEntity, serializer.MakeErrResponse(errno.ErrFileTypeInvalid))
		return
	}

	//在本地保存头像图片,并检查目录是否存在
	dir := "./upload/avator" + util.GetDay()
	if _, err := os.Stat(dir); err != nil {

		//生成图片保存目录
		if err := os.MkdirAll(dir, 0666); err != nil {
			c.JSON(http.StatusInternalServerError, serializer.MakeErrResponse(errno.ErrSystem))
			return
		}
	}

	//生成图片名称
	avatorName := util.RandStringByte(3) + strconv.FormatInt(util.GetUnix(), 10)
	avator.Filename = avatorName + avatorExt
	dst := path.Join(dir, avator.Filename)

	//上传图片至指定路径
	if err := c.SaveUploadedFile(avator, dst); err != nil {
		c.JSON(http.StatusBadRequest, serializer.MakeErrResponse(errno.ErrSaveFileFail))
		return
	}

	//从上下文中拿出鉴权用户id
	uid := c.GetUint("uid")
	res := UploadAvatorService.UpLoadAvator(uid, dst)
	c.JSON(http.StatusOK, res)
}
