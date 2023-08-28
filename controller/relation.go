package controller

import (
	"github.com/RaymondCode/simple-demo/models"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserListResponse struct {
	models.Response
	UserList []models.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	toUserId := c.Query("to_user_id")
	actionType := c.Query("action_type")

	rawId, ok := c.Get("user_id")
	if !ok {
		models.Fail(c, 1, "token解析出错")
		return
	}
	userId, ok := rawId.(int64) //保证id是string
	if !ok {
		models.Fail(c, 1, "user_id不是int64类型")
		return
	}

	res := &service.Response{}
	err := res.RelationAction(userId, utils.StringToInt(toUserId), actionType)
	if err != nil {
		models.Fail(c, 1, err.Error())
		return
	}
	res.StatusCode = 0
	res.StatusMsg = "ok"
	c.JSON(http.StatusOK, res)
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	rawId, ok := c.Get("user_id")
	if !ok {
		models.Fail(c, 1, "token解析出错")
		return
	}
	userId, ok := rawId.(int64) //保证id是string
	if !ok {
		models.Fail(c, 1, "user_id不是int64类型")
		return
	}

	res := &service.FollowResponse{}
	err := res.GetFollowList(userId, "follow_userId", "follower_userId")
	if err != nil {
		models.Fail(c, 1, err.Error())
		return
	}
	res.StatusCode = 0
	res.StatusMsg = "ok"
	c.JSON(http.StatusOK, res)
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	rawId, ok := c.Get("user_id")
	if !ok {
		models.Fail(c, 1, "token解析出错")
		return
	}
	userId, ok := rawId.(int64) //保证id是string
	if !ok {
		models.Fail(c, 1, "user_id不是int64类型")
		return
	}

	res := &service.FollowResponse{}
	err := res.GetFollowList(userId, "follower_userId", "follow_userId")
	if err != nil {
		models.Fail(c, 1, err.Error())
		return
	}
	res.StatusCode = 0
	res.StatusMsg = "ok"
	c.JSON(http.StatusOK, res)
}

// FriendList all users have same friend list
func FriendList(c *gin.Context) {
	rawId, ok := c.Get("user_id")
	if !ok {
		models.Fail(c, 1, "token解析出错")
		return
	}
	userId, ok := rawId.(int64) //保证id是string
	if !ok {
		models.Fail(c, 1, "user_id不是int64类型")
		return
	}

	res := &service.FollowResponse{}
	err := res.GetFriendList(userId)
	if err != nil {
		models.Fail(c, 1, err.Error())
		return
	}
	res.StatusCode = 0
	res.StatusMsg = "ok"
	c.JSON(http.StatusOK, res)
}
