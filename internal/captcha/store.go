package captcha

import (
	"context"
	"ruoyi-go/constant"
	"ruoyi-go/framework/dal"
	"time"
)

type store struct{}


func NewStore() *store {
	return &store{}
}

// 设置验证码
func (s *store) Set(id string, value string) error {

	key := constant.SystemCaptcha + id

	return dal.Redis.Set(context.Background(), key, value, time.Minute*2).Err()
}

// 获取验证码
func (s *store) Get(id string, clear bool) string {

	key := constant.SystemCaptcha + id

	captcha := dal.Redis.Get(context.Background(), key).String()

	if clear {
		dal.Redis.Del(context.Background(), key)
	}

	return captcha
}

// 验证验证码
func (s *store) Verify(id, answer string, clear bool) bool {
	return s.Get(id, clear) == answer
}
