package controller

import (
	"net/http"
	"strconv"

	"github.com/RaymondCode/simple-demo/models"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	// token := c.Query("token")

	// if _, exist := usersLoginInfo[token]; exist {
	// 	c.JSON(http.StatusOK, models.Response{StatusCode: 0, StatusMsg: "ok"})
	// } else {
	// 	c.JSON(http.StatusOK, models.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	// }

	rawID, ok := c.Get("user_id")
	if !ok {
		models.Fail(c, 1, "tokne解析出错")
		return
	}
	userID, ok := rawID.(int64)
	if !ok {
		models.Fail(c, 1, "用户名ID解析出错")
		return
	}
	actionType := c.Query("action_type")
	videoIDStr := c.Query("video_id")
	videoID, err := strconv.ParseInt(videoIDStr, 10, 64)
	if err != nil {
		models.Fail(c, 1, err.Error())
		return
	}
	if actionType == "1" { //点赞
		if err := service.Favorite(userID, videoID); err != nil {
			models.Fail(c, 1, err.Error())
			return
		}
		c.JSON(http.StatusOK, models.Response{StatusCode: 0, StatusMsg: "点赞成功"})
	} else if actionType == "2" { // 取消点赞
		if err := service.Unfavorite(userID, videoID); err != nil {
			models.Fail(c, 1, err.Error())
			return
		}
		c.JSON(http.StatusOK, models.Response{StatusCode: 0, StatusMsg: "取消点赞成功"})
	} else {
		models.Fail(c, 1, "action_type出错")
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	//自己的userId
	myUserId, ok := c.Get("user_id")
	if !ok {
		models.Fail(c, 1, "tokne解析出错")
		return
	}

	if myUserId, ok = myUserId.(int64); !ok {
		models.Fail(c, 1, "用户名ID解析出错")
		return
	}
	videoList, err := service.FavoriteList(myUserId.(int64))
	if err != nil {
		models.Fail(c, 1, "视频错误")
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response:  models.Response{StatusCode: 0, StatusMsg: "ok"},
		VideoList: videoList,
	})
}
