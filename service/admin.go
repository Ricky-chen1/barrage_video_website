package service

import (
	"barrage_video_website/dao"
	"barrage_video_website/model"
	"barrage_video_website/pkg/errno"
	"barrage_video_website/pkg/util"
	"barrage_video_website/serializer"
)

type AdminLoginService struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type AdminAddService struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

// 添加管理员
func (service *AdminAddService) Add() serializer.Response {
	if _, ok := dao.NewAdmin().IsAdminExist(service.Email); ok {
		return serializer.MakeErrResponse(errno.ErrAdminExist)
	}

	//密码加密
	password, err := util.SetPassword(service.Password)
	if err != nil {
		return serializer.MakeErrResponse(errno.ErrEncryptPassWordFail)
	}

	//调用dao层方法创建记录
	if err := dao.NewAdmin().Create(map[string]interface{}{
		"email":           service.Email,
		"password_digest": password,
		"authority":       1,
	}); err != nil {
		return serializer.MakeErrResponse(errno.ErrCreateUserFail)
	}

	return serializer.Response{
		Status: errno.Success.Code,
		Msg:    errno.Success.Msg,
	}
}

// 管理员登录(邮箱登录)
func (service *AdminLoginService) Login() serializer.Response {
	var admin *model.Admin
	admin, ok := dao.NewAdmin().IsAdminExist(service.Email)
	if !ok {
		return serializer.MakeErrResponse(errno.ErrAdminNoExist)
	}

	if err := util.ComparePassword(service.Password, admin.PasswordDigest); err != nil {
		return serializer.MakeErrResponse(errno.ErrPassWordWrong)
	}

	token, err := util.GenerateAdminToken(admin.ID)
	if err != nil {
		return serializer.MakeErrResponse(errno.ErrTokenGenerateFail)
	}

	return serializer.Response{
		Status: errno.Success.Code,
		Msg:    errno.Success.Msg,
		Data: serializer.AdminTokenData{
			Token: token,
			Admin: *serializer.BuildAdmin(admin),
		},
	}
}
