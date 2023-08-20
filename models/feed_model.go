package models

import (
	"errors"
	"time"
)

// QueryVideoListByLimitAndTime  返回按投稿时间倒序的视频列表，并限制为最多limit个
func QueryVideoListByLimitAndTime(limit int, latestTime time.Time, videoList *[]*Video) error {
	if videoList == nil {
		return errors.New("QueryVideoListByLimit videoList 空指针")
	}
	return DB.Model(&Video{}).Where("created_at<?", latestTime).
		Order("created_at DESC").Limit(limit).
		Preload("Author").Find(videoList).Error
}

func QueryIsFavorite(videoId int64, userId int64) bool {
	var favorite Favorite
	res := DB.Table("favorites").Where("video_id = ? and user_id = ?", videoId, userId).First(&favorite)
	if res.Error != nil || res.RowsAffected == 0 {
		return false
	}
	return true
}

func QueryFavoriteByUserID(UserId int64) ([]Favorite, error) {
	var favorite []Favorite
	res := DB.Table("favorites").Where("user_id = ?", UserId).Find(&favorite)
	return favorite, res.Error
}
