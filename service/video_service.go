package service

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/config"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/RaymondCode/simple-demo/models"
	"github.com/gin-gonic/gin"
)

func PublishVideo(c *gin.Context, authorId int64, data *multipart.FileHeader, title string) {
	// 存储视频
	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", authorId, filename)
	saveFile := filepath.Join("./public/videos", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 视频信息存入数据库
	video := &models.Video{
		Title:    title,
		AuthorID: authorId,
		PlayUrl:  config.Config.Url + "/static/videos/" + finalName,
		CoverUrl: "",
	}
	models.AddVideo(video)
}

// 发布列表
func PublishList(userId int64) (videoList []models.Video, err error) {
	videos, err := models.QueryPublishByUserId(userId)
	if err != nil {
		return nil, err
	}
	videoIds := make([]int64, len(videos))
	for i, videos := range videos {
		videoIds[i] = videos.Id
		video := models.Video{}
		result := models.DB.Model(&models.Video{}).Where("Id = ?", videoIds[i]).Find(&video)
		if result == nil {
			return nil, result.Error
		}
		//fmt.Println(videoIds[i])
		author := models.User{}
		res := models.DB.Model(&models.User{}).Where("Id = ?", video.AuthorID).Find(&author)
		if res == nil {
			return nil, res.Error
		}
		video.Author = author
		videoList = append(videoList, video)
	}
	return videoList, err
}
