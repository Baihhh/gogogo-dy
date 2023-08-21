package controller

import (
	"github.com/RaymondCode/simple-demo/models"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, models.Response{StatusCode: 0, StatusMsg: "ok"})
	} else {
		c.JSON(http.StatusOK, models.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
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
