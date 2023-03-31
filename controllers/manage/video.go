package manage

import (
	"barrage_video_website/pkg/errno"
	"barrage_video_website/serializer"
	"barrage_video_website/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VideoReview(c *gin.Context) {
	var videoReview service.VideoReviewService
	if err := c.ShouldBind(&videoReview); err != nil {
		c.JSON(http.StatusBadRequest, serializer.MakeErrResponse(errno.ErrQueryPramsInvalid))
		return
	}
	res := videoReview.Review()
	c.JSON(http.StatusOK, res)
}
