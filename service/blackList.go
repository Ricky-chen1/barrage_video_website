package service

import (
	"barrage_video_website/cache"
	"barrage_video_website/dao"
	"barrage_video_website/pkg/errno"
	"barrage_video_website/pkg/util"
	"barrage_video_website/serializer"
)

type UserBlacklistService struct {
	Uid uint `json:"uid" form:"uid"`
}

type UserWhitelistService struct {
	Uid uint `json:"uid" form:"uid"`
}

func (service *UserBlacklistService) Blacklist() serializer.Response {
	if !dao.NewUser().IsUserExist(service.Uid) {
		return serializer.MakeErrResponse(errno.ErrNotExistUser)
	}

	user, err := dao.NewUser().QueryUser(map[string]interface{}{
		"id": service.Uid,
	})
	if err != nil {
		return serializer.MakeErrResponse(errno.ErrQueryUserInfoFail)
	}

	//判断用户是否已拉黑
	if user.IsBlackList() {
		return serializer.MakeErrResponse(errno.ErrIsUserBlackList)
	}

	//将用户添加至黑名单
	user.BlacklistAdd()

	cache.RedisNewClient.Set(cache.Ctx, cache.TokenKey(service.Uid), util.GetUnix(), 0)

	return serializer.Response{
		Status: errno.Success.Code,
		Msg:    errno.Success.Msg,
	}
}

func (service *UserWhitelistService) Whitelist() serializer.Response {
	if !dao.NewUser().IsUserExist(service.Uid) {
		return serializer.MakeErrResponse(errno.ErrNotExistUser)
	}

	user, err := dao.NewUser().QueryUser(map[string]interface{}{
		"id": service.Uid,
	})
	if err != nil {
		return serializer.MakeErrResponse(errno.ErrQueryUserInfoFail)
	}

	//将用户从黑名单中删除
	user.BlacklistDelete()

	return serializer.Response{
		Status: errno.Success.Code,
		Msg:    errno.Success.Msg,
	}
}
