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

// 比较工具
// 检查元素item是否存在于切片slice中
// 如果存在，返回true；如果不存在，返回false
func Contains[T comparable](slice []T, item T) bool {

	for _, value := range slice {
		if value == item {
			return true
		}
	}

	return false
}

// 过滤器
// 条件函数返回true，元素会被包含在结果中
func Filter[T interface{}](slice []T, condition func(T) bool) (result []T) {

	for _, value := range slice {
		if condition(value) {
			result = append(result, value)
		}
	}

	return result
}

// 脱敏工具
func Desensitize(content string, start, end int) string {

	if start < 0 || end < 0 || start > end {
		return content
	}

	var contentRune []rune

	for key, value := range content {
		if key >= start && key <= end {
			contentRune = append(contentRune, '*')
		} else {
			contentRune = append(contentRune, value)
		}
	}

	return string(contentRune)
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