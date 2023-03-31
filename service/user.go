package service

import (
	"barrage_video_website/cache"
	"barrage_video_website/dao"
	"barrage_video_website/model"
	"barrage_video_website/pkg/errno"
	"barrage_video_website/pkg/util"
	"barrage_video_website/serializer"
	"errors"
	"mime/multipart"

	"gorm.io/gorm"
)

type UserRegisterService struct {
	Username string `json:"username" form:"username" binding:"required,min=3,max=16"`
	Password string `json:"password" form:"password" binding:"required,min=5,max=16"`
	Email    string `json:"email" form:"email" `
}

type UserLoginService struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type UserShowService struct{}

type UserUpdateNameService struct {
	NewUsername string `json:"new_username" form:"new_username" binding:"required,min=3,max=16"`
}

type UserUpdateEmailService struct {
	NewEmail string `json:"new_email" form:"new_email" binding:"required"`
}

type UserModifyPasswordService struct {
	NewPassword string `json:"new_password" form:"new_password"`
}

type UserAvatorUploadService struct {
	Avator *multipart.FileHeader `form:"avator"`
}

// 用户注册
func (service *UserRegisterService) Register() serializer.Response {
	//判断用户名是否已被注册
	_, err := dao.NewUser().QueryUser(map[string]interface{}{
		"username": service.Username,
	})
	if err == nil {
		return serializer.MakeErrResponse(errno.ErrExistUser)
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return serializer.MakeErrResponse(errno.ErrQueryUserInfoFail)
	}

	//设置密码(加密)
	passwordDigest, err := util.SetPassword(service.Password)
	if err != nil {
		return serializer.MakeErrResponse(errno.ErrEncryptPassWordFail)
	}

	newUser := model.User{
		Username:       service.Username,
		Email:          service.Email,
		PasswordDigest: passwordDigest,
	}

	//用户注册
	user, err := dao.NewUser().Create(newUser)
	if err != nil {
		return serializer.MakeErrResponse(errno.ErrCreateUserFail)
	}

	//注册成功后创建默认收藏夹
	collection := model.Collection{
		Uid:  user.ID,
		Name: "默认收藏夹",
	}
	if err := dao.NewCollection().Create(collection); err != nil {
		return serializer.MakeErrResponse(errno.ErrCollectionCreateFail)
	}

	return serializer.Response{
		Status: errno.Success.Code,
		Msg:    errno.Success.Msg,
	}
}

// 用户登录
func (service *UserLoginService) Login() serializer.Response {
	var (
		user *model.User
		err  error
	)

	//查询用户是否存在
	user, err = dao.NewUser().QueryUser(map[string]interface{}{
		"username": service.Username,
	})
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return serializer.MakeErrResponse(errno.ErrNotExistUser)
	}
	if err != nil {
		return serializer.MakeErrResponse(errno.ErrQueryUserInfoFail)
	}

	//用户是否被拉黑
	if user.IsBlackList() {
		return serializer.MakeErrResponse(errno.ErrIsUserBlackList)
	}

	//密码校验
	if err := util.ComparePassword(service.Password, user.PasswordDigest); err != nil {
		return serializer.MakeErrResponse(errno.ErrPassWordWrong)
	}

	//token签发
	token, err := util.GenerateToken(user.ID, user.Username, user.Email)
	if err != nil {
		return serializer.MakeErrResponse(errno.ErrTokenGenerateFail)
	}

	return serializer.Response{
		Status: errno.Success.Code,
		Msg:    errno.Success.Msg,
		Data: serializer.TokenData{
			User:  *serializer.BuildUser(user),
			Token: token,
		},
	}
}

// 展示个人资料
func (service *UserShowService) Show(uid uint) serializer.Response {
	var (
		user *model.User
		err  error
	)

	//查询用户信息
	if user, err = dao.NewUser().QueryUser(map[string]interface{}{
		"id": uid,
	}); err == gorm.ErrRecordNotFound {
		return serializer.MakeErrResponse(errno.ErrNotExistUser)
	}
	if err != nil {
		return serializer.MakeErrResponse(errno.ErrQueryUserInfoFail)
	}

	//展示用户信息
	return serializer.Response{
		Status: errno.Success.Code,
		Msg:    errno.Success.Msg,
		Data:   serializer.BuildUser(user),
	}
}

// 更改用户名
func (service *UserUpdateNameService) UpdateName(uid uint) serializer.Response {
	if err := dao.NewUser().Update(uid, map[string]interface{}{
		"username": service.NewUsername,
	}); err != nil {
		return serializer.MakeErrResponse(errno.ErrUpdateUserFail)
	}

	return serializer.Response{
		Status: errno.Success.Code,
		Msg:    errno.Success.Msg,
		Data:   service.NewUsername,
	}
}

// 更改绑定邮箱
func (service *UserUpdateEmailService) UpdateEmail(uid uint) serializer.Response {
	if err := dao.NewUser().Update(uid, map[string]interface{}{
		"email": service.NewEmail,
	}); err != nil {
		return serializer.MakeErrResponse(errno.ErrUpdateUserFail)
	}

	return serializer.Response{
		Status: errno.Success.Code,
		Msg:    errno.Success.Msg,
		Data:   service.NewEmail,
	}
}

// 修改用户密码(是否要重新登录?)
func (service *UserModifyPasswordService) Modify(uid uint) serializer.Response {
	var err error

	//找到要更改的用户
	_, err = dao.NewUser().QueryUser(map[string]interface{}{
		"id": uid,
	})
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return serializer.MakeErrResponse(errno.ErrNotExistUser)
	}
	if err != nil {
		return serializer.MakeErrResponse(errno.ErrQueryUserInfoFail)
	}

	//修改密码并加密
	new_password, err := util.SetPassword(service.NewPassword)
	if err != nil {
		return serializer.MakeErrResponse(errno.ErrModifyPasswordFail)
	}
	if err := dao.NewUser().Update(uid, map[string]interface{}{
		"password_digest": new_password,
	}); err != nil {
		return serializer.MakeErrResponse(errno.ErrModifyPasswordFail)
	}

	//更新缓存中token过期时间
	cache.RedisNewClient.Set(cache.Ctx, cache.TokenKey(uid), util.GetUnix(), 0)

	return serializer.Response{
		Status: errno.Success.Code,
		Msg:    errno.Success.Msg,
	}
}

// 更改用户头像
func (service *UserAvatorUploadService) UpLoadAvator(uid uint, dst string) serializer.Response {
	if err := dao.NewUser().Update(uid, map[string]interface{}{
		"avator": dst,
	}); err != nil {
		return serializer.MakeErrResponse(errno.ErrUpdateAvatorFail)
	}

	return serializer.Response{
		Status: errno.Success.Code,
		Msg:    errno.Success.Msg,
	}
}
