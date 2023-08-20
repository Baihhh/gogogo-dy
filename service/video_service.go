package service

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/RaymondCode/simple-demo/config"
	"github.com/RaymondCode/simple-demo/utils"

	"github.com/RaymondCode/simple-demo/models"
	"github.com/gin-gonic/gin"
)

func PublishVideo(c *gin.Context, authorId int64, data *multipart.FileHeader, title string) error {
	// 存储视频
	filename := filepath.Base(data.Filename)
	if !utils.ValidateVideoFile(filepath.Ext(filename)) {
		return errors.New("视频格式不支持")
	}
	finalName := fmt.Sprintf("%d_%s", authorId, filename)
	saveFile := filepath.Join("./public/videos", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		return err
	}
	// 拼接封面名
	coverName := finalName[0 : len(finalName)-len(filepath.Ext(filename))] // 不包含后缀名
	coverName = coverName + ".jpg"
	coverUrl := filepath.Join("./public/video_covers", coverName)

	workPath, _ := os.Getwd()
	cmd := exec.Command(workPath+"./ffmpeg", "-i", saveFile, "-ss", "00:00:00.001", "-vframes", "1", coverUrl)
	err := cmd.Run()
	if err != nil {
		return err
	}

	// 视频信息存入数据库
	video := &models.Video{
		Title:    title,
		AuthorID: authorId,
		PlayUrl:  config.Config.Url + "/static/videos/" + finalName,
		CoverUrl: config.Config.Url + "/static/video_covers/" + coverName,
	}
	models.AddVideo(video)
	return nil
}
