package controller

import (
	"fmt"
	"net/http"

	"github.com/RaymondCode/simple-demo/models"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	// 不会报错
	token, ok := c.GetQuery("token")
	res := &service.FeedResponse{}
	res.VideoList = []*models.Video{}
	//没有token
	if !ok || token == "" {
		err := res.DoNoToken(c)
		if err != nil {
			models.Fail(c, 1, err.Error())
			return
		}

	} else {
		//有token
		err := res.DoHasToken(token, c)
		if err != nil {
			models.Fail(c, 1, err.Error())
			return
		}
	}
	res.StatusCode = 0
	res.StatusMsg = "ok"
	fmt.Print(res)
	c.JSON(http.StatusOK, res)

}
