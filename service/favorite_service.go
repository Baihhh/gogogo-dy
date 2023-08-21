package service

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/models"
)

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
		videoIds[i] = favorites.VideoId
		video := models.Video{}
		result := models.DB.Model(&models.Video{}).Preload("User").Where("id = ?", videoIds[i]).Find(&video)
		if result == nil {
			return nil, result.Error
		}
		fmt.Println(videoIds[i])
		videoList = append(videoList, video)
	}
	return videoList, err
}
