package models

import (
	"log"

	"github.com/RaymondCode/simple-demo/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() {
	var err error
	dataSource := config.Config.MySql.DbUser + ":" + config.Config.MySql.DbPwd + "@tcp(" + config.Config.MySql.DbHost + ":" + config.Config.MySql.DbPort + ")/" + config.Config.MySql.DbName + "?charset=" + config.Config.MySql.DbCharset + "&parseTime=true"
	DB, err = gorm.Open(mysql.Open(dataSource), &gorm.Config{
		PrepareStmt:            true, //缓存预编译命令
		SkipDefaultTransaction: true, //禁用默认事务操作
	})
	if err != nil {
		log.Fatalf("Fail to connect DB: %v", err)
	}

	// 数据库表动态迁移
	DB.AutoMigrate(&User{}, &Follow{}, &Video{}, &Comment{}, &Favorite{}, &Message{})
}
