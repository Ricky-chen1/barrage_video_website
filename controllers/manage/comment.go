package manage

import (
	"barrage_video_website/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 管理员删除视频评论
func CommentDelete(c *gin.Context) {
	var commentDelete service.CommentDeleteService

	cid, _ := strconv.ParseUint(c.Param("cid"), 10, 64)
	res := commentDelete.Delete(uint(cid))
	c.JSON(http.StatusOK, res)
}
