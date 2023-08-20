package controller

import (
	"net/http"

	"github.com/RaymondCode/simple-demo/models"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	models.Response
	VideoList []models.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {

	// 获取发布者信息
	rawId, ok := c.Get("user_id")
	if !ok {
		models.Fail(c, 1, "token解析出错")
		return
	}
	authorId, ok := rawId.(int64) //保证id是int64
	if !ok {
		models.Fail(c, 1, "user_id不是int64类型")
		return
	}
	data, err := c.FormFile("data")
	if err != nil {
		models.Fail(c, 1, err.Error())
		return
	}
	title := c.PostForm("title")
	if err := service.PublishVideo(c, authorId, data, title); err != nil {
		models.Fail(c, 1, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.Response{
		StatusCode: 0,
		StatusMsg:  title + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: models.Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
