package service

import (
	"errors"
	"strconv"
	"time"

	"github.com/RaymondCode/simple-demo/middleware"
	"github.com/RaymondCode/simple-demo/models"
	"github.com/gin-gonic/gin"
)

type FeedResponse models.FeedResponse

const (
	MaxVideoNum = 10
)

func (f *FeedResponse) DoNoToken(c *gin.Context) error {
	lastTime := getLastTime(c)
	err := models.QueryVideoListByLimitAndTime(MaxVideoNum, lastTime, &f.VideoList)
	if err != nil {
		return err
	}
	if len(f.VideoList) != 0 {
		lastTime = f.VideoList[(len(f.VideoList) - 1)].CreatedAt
	} else {
		lastTime = time.Now()
	}

	f.NextTime = lastTime.Unix() * 1e3
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
		err = fillFollowAndFavorite(claim.UserId, &f.VideoList)
		if err != nil {
			return err
		}
		var latestTime time.Time
		if len(f.VideoList) != 0 {
			latestTime = f.VideoList[(len(f.VideoList) - 1)].CreatedAt
		} else {
			latestTime = time.Now()
		}
		f.NextTime = latestTime.Unix() * 1e3
		return nil
	}
	return nil
}

func fillFollowAndFavorite(userId int64, videos *[]*models.Video) error {
	size := len(*videos)
	// if videos == nil || size == 0 {
	// 	return nil, errors.New("没有可以播放的视频")
	// }

	for i := 0; i < size; i++ {
		(*videos)[i].Author.IsFollow = models.QueryIsFollow((*videos)[i].Author.Id, userId)
		//填充有登录信息的点赞状态
		if userId > 0 {
			(*videos)[i].IsFavorite = models.QueryIsFavorite((*videos)[i].Id, userId)
		}
	}
	return nil
}

func getLastTime(c *gin.Context) (latestTime time.Time) {
	rawTimestamp, ok := c.GetQuery("latest_time")
	if ok {
		intTime, err := strconv.ParseInt(rawTimestamp, 10, 64)
		if err == nil && intTime != 0 {
			latestTime = time.Unix(intTime/1e3, 0)
			return latestTime
		}
	}
	latestTime = time.Now()
	return latestTime
}
