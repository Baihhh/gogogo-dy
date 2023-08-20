package models

import (
	"context"

	"github.com/RaymondCode/simple-demo/config"
	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client
var Ctx = context.Background()

func InitRedis() {
	//初始化redis，连接地址和端口，密码，数据库名称
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Addr,
		Password: config.Config.Redis.Password,
		DB:       config.Config.Redis.DB,
	})
}
