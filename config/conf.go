package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"path/filepath"
)

var Cfg *ini.File
var Config Conf

type Conf struct {
	Url   string
	Port  string
	MySql Mysql
	Jwt   Jwt
}

type Mysql struct {
	DbUser    string `json:db_user`
	DbName    string `json:db_name`
	DbPwd     string `json:db_pwd`
	DbHost    string `json:db_host`
	DbPort    string `json:db_port`
	DbCharset string `json:db_charset`
}

type Jwt struct {
	Secret string `mapstructure:"secret" json:"secret" yaml:"secret"`
	JwtTtl int64  `mapstructure:"jwt_ttl" json:"jwt_ttl" yaml:"jwt_ttl"` // token 有效期（秒）
}

func InitConfig() {
	var err error
	// 获取当前工作目录
	currentDir, _ := os.Getwd()

	// 构建配置文件的绝对路径
	configPath := filepath.Join(currentDir, "config/config.ini")
	Cfg, err = ini.Load(configPath)
	if err != nil {
		log.Fatalf("Fail to parse 'config.ini': %v", err)
	}

	loadApp()
	loadMysql()
}

func loadApp() {
	Config.Port = getConfig("app", "port")
	Config.Url = fmt.Sprintf("%s:%s", getConfig("app", "url"), Config.Port)
	Config.Jwt.Secret = getConfig("jwt", "secret")
	Config.Jwt.JwtTtl, _ = Cfg.Section("jwt").Key("jwt_ttl").Int64()
}

func loadMysql() {
	Config.MySql = Mysql{
		DbName:    getConfig("mysql", "db_name"),
		DbUser:    getConfig("mysql", "db_user"),
		DbPwd:     getConfig("mysql", "db_pwd"),
		DbHost:    getConfig("mysql", "db_host"),
		DbPort:    getConfig("mysql", "db_port"),
		DbCharset: getConfig("mysql", "db_charset"),
	}
}

func getConfig(name string, key string) string {
	return Cfg.Section(name).Key(key).String()
}
