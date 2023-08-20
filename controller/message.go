package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/models"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

var tempChat = map[string][]models.Message{}

type ChatResponse struct {
	models.Response
	MessageList []models.Message `json:"message_list"`
}

// MessageAction no practical effect, just check if token is valid
func MessageAction(c *gin.Context) {
	toUserId := c.Query("to_user_id")
	actionType := c.Query("action_type")
	content := c.Query("content")

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

	err := service.MessageSend(userId, utils.StringToInt(toUserId), content, actionType)
	if err != nil {
		models.Fail(c, 1, err.Error())
		return
	}
	c.JSON(http.StatusOK, models.Response{StatusCode: 0, StatusMsg: "ok"})

}

// MessageChat all users have same follow list
func MessageChat(c *gin.Context) {
	toUserId := c.Query("to_user_id")

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
	res := &service.MessageResponse{}
	res.MessageList = []*models.Message{}
	err := res.GetMessage(userId, utils.StringToInt(toUserId))
	if err != nil {
		models.Fail(c, 1, err.Error())
		return
	}
	res.StatusCode = 0
	res.StatusMsg = "ok"
	c.JSON(http.StatusOK, res)

	//if user, exist := usersLoginInfo[token]; exist {
	//	userIdB, _ := strconv.Atoi(toUserId)
	//	chatKey := genChatKey(user.Id, int64(userIdB))
	//
	//	c.JSON(http.StatusOK, ChatResponse{Response: models.Response{StatusCode: 0}, MessageList: tempChat[chatKey]})
	//} else {
	//	c.JSON(http.StatusOK, models.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//}
}

func genChatKey(userIdA int64, userIdB int64) string {
	if userIdA > userIdB {
		return fmt.Sprintf("%d_%d", userIdB, userIdA)
	}
	return fmt.Sprintf("%d_%d", userIdA, userIdB)
}
