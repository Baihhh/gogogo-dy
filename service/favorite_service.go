package service

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/models"
)

// 点赞
func Favorite(userID int64, videoID int64) error {
	err := models.AddFavorite(userID, videoID)
	return err
}

// 取消点赞
func Unfavorite(userID int64, videoID int64) error {
	err := models.DelFavorite(userID, videoID)
	return err
}

// 点赞列表
func FavoriteList(userId int64) (videoList []models.Video, err error) {
	//var res []models.Video
	favorites, err := models.QueryFavoriteByUserID(userId)
	if err != nil {
		return nil, err
	}
	if len(favorites) == 0 {
		return []models.Video{}, nil
	}
	videoIds := make([]int64, len(favorites))
	for i, favorites := range favorites {
		videoIds[i] = favorites.VideoID
		video := models.Video{}
		result := models.DB.Model(&models.Video{}).Preload("Author").Where("id = ?", videoIds[i]).Find(&video)
		if result == nil {
			return nil, result.Error
		}
		fmt.Println(videoIds[i])
		videoList = append(videoList, video)
	}
	return videoList, err
}
