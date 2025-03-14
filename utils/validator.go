package utils

import (
	"errors"
	"regexp"
	"strings"

	"github.com/RaymondCode/simple-demo/config"
	"github.com/RaymondCode/simple-demo/models"
	"github.com/go-playground/validator/v10"
)

const (
	MaxUsernameLength = 100
	MaxPasswordLength = 20
	MinPasswordLength = 8
)

// 校验视频文件
func ValidateVideoFile(fileExt string) bool {
	extensions := strings.Split(config.Config.FileExt.VideoExt, "/")
	for _, allowedExt := range extensions {
		if fileExt == allowedExt {
			return true
		}
	}
	return false
}

// 校验图像文件
func ValidateImageFile(fileExt string) bool {
	extensions := strings.Split(config.Config.FileExt.ImageExt, "/")
	for _, allowedExt := range extensions {
		if fileExt == allowedExt {
			return true
		}
	}
	return false
}

// ValidateMobile 校验手机号
func ValidateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	ok, _ := regexp.MatchString(`^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$`, mobile)
	if !ok {
		return false
	}
	return true
}

// ValidateRegister
func ValidateRegister(username string, password string, key string) error {
	if err := ValidateNameAndPwd(username, password); err != nil {
		return err
	}
	if hasSpecialChar(username) {
		return errors.New("用户名不能包含特殊字符")
	}
	if models.IsUserExistByUsername(username) && key == "register" {
		return errors.New("用户名已经存在")
	}
	if !models.IsUserExistByUsername(username) && key == "login" {
		return errors.New("该用户尚未注册")
	}
	return nil
}

// ValidateNameAndPwd
func ValidateNameAndPwd(username string, password string) error {
	if username == "" {
		return errors.New("用户名不能为空")
	}
	if len(username) > MaxUsernameLength {
		return errors.New("用户名长度超出限制")
	}
	if password == "" {
		return errors.New("密码不能为空")
	}
	return nil
}

// ValidateActionType
func ValidateActionType(actionType string) error {
	if actionType != "1" && actionType != "2" {
		return errors.New("错误的操作类型")
	}
	return nil
}

func hasSpecialChar(username string) bool {
	reg := regexp.MustCompile(`[#%￥&《》<>?:'"{})(*^$!~]`)
	return reg.MatchString(username)
}
