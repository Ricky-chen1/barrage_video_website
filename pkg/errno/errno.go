package errno

type Errno struct {
	Code    int
	Message string
}

type Ok struct {
	Code int
	Msg  string
}

var Success = &Ok{Code: 200, Msg: "success"}

var (
	ErrQueryPramsInvalid = &Errno{Code: 40001, Message: "请求参数不合法"}
	// 数据库相关 500 开头
	ErrDataBase          = &Errno{Code: 50001, Message: "数据库错误"}
	ErrQueryUserInfoFail = &Errno{Code: 50002, Message: "查询用户信息错误"}
	ErrCreateUserFail    = &Errno{Code: 50003, Message: "创建用户信息失败"}
	ErrUpdateUserFail    = &Errno{Code: 50104, Message: "修改用户信息失败"}

	// Token、权限相关 401 开头
	ErrTokenExpired      = &Errno{Code: 40101, Message: "Token 已过期,请重新登录"}
	ErrTokenGenerateFail = &Errno{Code: 40102, Message: "Token 签发失败"}
	ErrNoToken           = &Errno{Code: 40103, Message: "No Token"}
	ErrTokenParseFail    = &Errno{Code: 40104, Message: "没有权限执行操作"}

	//用户相关 101 开头
	ErrExistUser           = &Errno{Code: 10101, Message: "用户已存在"}
	ErrNotExistUser        = &Errno{Code: 10102, Message: "用户不存在"}
	ErrAdminNoExist        = &Errno{Code: 10103, Message: "该管理员不存在"}
	ErrAdminExist          = &Errno{Code: 10104, Message: "该管理员已存在"}
	ErrPassWordWrong       = &Errno{Code: 10105, Message: "密码错误"}
	ErrEncryptPassWordFail = &Errno{Code: 10106, Message: "密码加密失败"}
	ErrModifyPasswordFail  = &Errno{Code: 10107, Message: "修改密码失败"}
	ErrUpdateAvatorFail    = &Errno{Code: 10108, Message: "更改头像失败"}
	ErrIsUserBlackList     = &Errno{Code: 10109, Message: "用户已被拉黑"}

	//视频相关 102 开头
	ErrVideoUpload           = &Errno{Code: 10201, Message: "视频上传失败"}
	ErrVideoNoExist          = &Errno{Code: 10202, Message: "视频不存在"}
	ErrCreateVideoRecordFail = &Errno{Code: 10203, Message: "数据库新增视频记录失败"}
	ErrQueryVideosFail       = &Errno{Code: 10204, Message: "查询视频信息失败"}
	ErrReviewVideosFail      = &Errno{Code: 10205, Message: "审核视频失败"}
	ErrRankVideosFail        = &Errno{Code: 10206, Message: "视频排行失败"}

	// 评论相关 103 开头
	ErrCommentNoExist       = &Errno{Code: 10301, Message: "评论已不存在"}
	ErrCommentAddFail       = &Errno{Code: 10302, Message: "新增评论失败"}
	ErrCommentDelFail       = &Errno{Code: 10303, Message: "删除评论失败"}
	ErrQueryCommentFail     = &Errno{Code: 10304, Message: "查询评论失败"}
	ErrQueryCommentListFail = &Errno{Code: 10304, Message: "查询评论列表失败"}
	ErrCommentReplyFail     = &Errno{Code: 10305, Message: "回复评论失败"}

	//文件相关104开头
	ErrGetFileFail     = &Errno{Code: 10401, Message: "获取文件失败"}
	ErrSaveFileFail    = &Errno{Code: 10402, Message: "保存文件失败"}
	ErrPathNoExist     = &Errno{Code: 10403, Message: "该目录不存在"}
	ErrSystem          = &Errno{Code: 10404, Message: "创建文件夹失败"}
	ErrFileTypeInvalid = &Errno{Code: 10405, Message: "文件类型不匹配"}

	//互动相关105开头
	ErrFavorExist           = &Errno{Code: 10501, Message: "不可重复点赞"}
	ErrFavorFail            = &Errno{Code: 10502, Message: "点赞失败"}
	ErrCollectExist         = &Errno{Code: 10503, Message: "不可重复收藏"}
	ErrCollectFail          = &Errno{Code: 10504, Message: "收藏失败"}
	ErrCollectionCreateFail = &Errno{Code: 10504, Message: "创建收藏夹失败"}
	ErrCollectionNoExist    = &Errno{Code: 10505, Message: "收藏夹不存在"}
	ErrForwardFail          = &Errno{Code: 10506, Message: "转发失败"}
	ErrSendBarrageFail      = &Errno{Code: 10507, Message: "发送弹幕失败"}
	ErrQueryBarrageListFail = &Errno{Code: 10508, Message: "获取弹幕列表失败"}

	//搜索相关
	ErrSearchVideoFail = &Errno{Code: 10601, Message: "搜索视频失败"}
	ErrNoSearchVideo   = &Errno{Code: 10602, Message: "未找到相关视频"}
)
