package routers

import (
	"barrage_video_website/controllers"
	"barrage_video_website/controllers/manage"
	"barrage_video_website/middleware"

	"github.com/gin-gonic/gin"
)

// 路由注册服务
func RoutersInit(httpPort string) *gin.Engine {
	r := gin.Default()
	r.POST("/user/register", controllers.UserRegister)
	r.POST("/user/login", controllers.UserLogin)
	r.POST("/admin/login", manage.AdminLogin)
	r.POST("/admin", manage.AdminAdd)
	//点击量排行榜
	r.GET("/rank/daily", controllers.DailyRank)
	//搜索视频
	r.GET("/search/videos", controllers.SearchVideo)
	//获取弹幕列表
	r.GET("/barrages", controllers.GetBarrageList)
	//获取评论列表
	r.GET("/comments", controllers.CommentListGet)

	adminAuthed := r.Group("/api/v1/admin")
	adminAuthed.Use(middleware.AdminJWT())
	{
		adminAuthed.PUT("/video", manage.VideoReview)
		adminAuthed.PUT("/blacklist/user", manage.UserBlacklist)
		adminAuthed.DELETE("/comment/:cid", manage.CommentDelete)
		adminAuthed.PUT("/whitelist/user", manage.UserWhitelist)
	}

	authed := r.Group("/api/v1")
	authed.Use(middleware.JWT())
	{
		u := authed.Group("/user")
		{
			u.GET("/show", controllers.UserShow)
			u.PUT("/name", controllers.UserUpdateName)
			u.PUT("/email", controllers.UserUpdateEmail)
			u.POST("/avator", controllers.UserAvatorUpload)
			u.PUT("/password", controllers.UserModifyPassword)
		}

		v := authed.Group("/video")
		{
			v.POST("", controllers.VideoUpload)
			v.GET("/show", controllers.ShowVideo)
		}

		c := authed.Group("/comment")
		{
			c.POST("", controllers.CommentAdd)
			c.POST("/reply", controllers.Reply)
		}

		i := authed.Group("")
		{
			i.POST("/favor", controllers.Favor)
			i.POST("/collect", controllers.Collect)
			i.POST("/forward", controllers.Forward)
			i.POST("collection", controllers.CreateCollection)
		}

		b := authed.Group("/barrage")
		{
			b.POST("", controllers.SendBarrage)
		}
	}

	r.Run(httpPort)
	return r
}
