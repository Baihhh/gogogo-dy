package main

import (
	"github.com/RaymondCode/simple-demo/config"
	"github.com/RaymondCode/simple-demo/models"
	"github.com/RaymondCode/simple-demo/router"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	go service.RunMessageServer()

	r := gin.Default()

	router.InitRouter(r)
	config.InitConfig()
	models.InitDb()

	if err := r.Run(":" + config.Config.Port); err != nil {
		log.Fatalf("启动服务失败")
	}
}
