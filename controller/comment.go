package controller

import (
	"net/http"
	"strconv"

	"github.com/RaymondCode/simple-demo/models"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

type CommentListResponse struct {
	models.Response
	CommentList []models.Comment `json:"comment_list"`
}

type CommentActionResponse struct {
	models.Response
	Comment models.Comment `json:"comment"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	rawID, ok := c.Get("user_id")
	if !ok {
		models.Fail(c, 1, "token解析出错")
		return
	}
	userID, ok := rawID.(int64) //保证id是int64
	if !ok {
		models.Fail(c, 1, "user_id不是int64类型")
		return
	}
	actionType := c.Query("action_type")

	// 添加评论
	if actionType == "1" {
		text := c.Query("comment_text")
		videoIDStr := c.Query("video_id")
		videoID, err := strconv.ParseInt(videoIDStr, 10, 64)
		if err != nil {
			models.Fail(c, 1, err.Error())
			return
		}
		comment := service.AddComment(userID, videoID, text)
		models.DB.Model(&models.Comment{}).Preload("User").First(comment, comment.Id) //加载评论发布者
		c.JSON(http.StatusOK, CommentActionResponse{Response: models.Response{StatusCode: 0, StatusMsg: "添加评论成功"},
			Comment: models.Comment{
				Id:        comment.Id,
				User:      comment.User,
				Content:   comment.Content,
				CreatedAt: comment.CreatedAt,
			}})
		return
	} else { //删除评论
		commentIDStr := c.Query("comment_id")
		commentID, err := strconv.ParseInt(commentIDStr, 10, 64)
		if err != nil {
			models.Fail(c, 1, err.Error())
			return
		}

		if err := service.DelComment(commentID); err != nil {
			models.Fail(c, 1, err.Error())
			return
		}
		c.JSON(http.StatusOK, models.Response{StatusCode: 0, StatusMsg: "删除评论成功"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	videoIdStr := c.Query("video_id")
	videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		models.Fail(c, 1, "视频不存在")
	}
	commentList, err := service.CommentList(videoId)
	if err != nil {
		models.Fail(c, 1, "评论加载异常")
	}
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    models.Response{StatusCode: 0, StatusMsg: "ok"},
		CommentList: commentList,
	})
}
