package utils

import (
	"errors"
	"gex-api/app/internal/jwt"
	"regexp"
	"time"
)

// 正则验证
// expr 正则表达式
// content 要验证的内容
func CheckRegex(expr, content string) bool {

	r, err := regexp.Compile(expr)
	if err != nil {
		return false
	}

	return r.MatchString(content)
}

// 数据是否存在于切片
// 存在返回true
// 不存在返回false
func Contains[T comparable](slice []T, item T) bool {

	for _, v := range slice {
		if v == item {
			return true
		}
	}

	return false
}

// 获取授权信息
func GetTokenPayload(token string) (int, error) {
	
	if token == "" {
		return 0, errors.New("未授权")
	}

	payload := jwt.Parse(token)
	if payload == nil || payload.Id <= 0 || time.Now().After(payload.Expire) {
		return 0, errors.New("授权过期")
	}

	return payload.Id, nil
}