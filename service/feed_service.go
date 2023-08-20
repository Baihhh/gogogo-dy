package service

import (
	"errors"
	"github.com/RaymondCode/simple-demo/middleware"
	"github.com/RaymondCode/simple-demo/models"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type FeedResponse models.FeedResponse

const (
	MaxVideoNum = 30
)

func (f *FeedResponse) DoNoToken(c *gin.Context) error {
	lastTime := getLastTime(c)
	err := models.QueryVideoListByLimitAndTime(MaxVideoNum, lastTime, &f.VideoList)
	if err != nil {
		return err
	}
	f.NextTime = lastTime.Unix()
	return nil
}

func (f *FeedResponse) DoHasToken(token string, c *gin.Context) error {
	if claim, ok := middleware.ParseToken(token); ok {
		//token超时
		if time.Now().Unix() > claim.ExpiresAt {
			return errors.New("token超时")
		}
		err := f.DoNoToken(c)
		if err != nil {
			return err
		}

		//如果用户为登录状态，则更新该视频是否被该用户点赞的状态
		latestTime, err := fillFollowAndFavorite(claim.UserId, &f.VideoList)
		if err != nil {
			return err
		}
		f.NextTime = latestTime.Unix()
		return nil
	}
	return nil
}

func fillFollowAndFavorite(userId int64, videos *[]*models.Video) (*time.Time, error) {
	size := len(*videos)
	if videos == nil || size == 0 {
		return nil, errors.New("没有可以播放的视频")
	}

	latestTime := (*videos)[size-1].CreatedAt //获取最近的投稿时间
	for i := 0; i < size; i++ {
		(*videos)[i].Author.IsFollow = models.QueryIsFollow((*videos)[i].Author.Id, userId)
		//填充有登录信息的点赞状态
		if userId > 0 {
			(*videos)[i].IsFavorite = models.QueryIsFavorite((*videos)[i].Id, userId)
		}
	}
	return &latestTime, nil
}

func getLastTime(c *gin.Context) (latestTime time.Time) {
	rawTimestamp, ok := c.GetQuery("latest_time")
	if ok {
		intTime, err := strconv.ParseInt(rawTimestamp, 10, 64)
		if err == nil && intTime != 0 {
			latestTime = time.Unix(0, intTime)
			return latestTime
		}
	}
	latestTime = time.Now()
	return latestTime
}
