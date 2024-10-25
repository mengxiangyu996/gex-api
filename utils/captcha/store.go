package captcha

import (
	"context"
	"ruoyi-go/framework/dal"
	"time"
)

type store struct{}

const SystemCaptcha = "system:captcha:" // 系统验证码，"system:captcha:uuid"

func NewStore() *store {
	return &store{}
}

// 设置验证码
func (s *store) Set(id string, value string) error {

	key := SystemCaptcha + id

	return dal.Redis.Set(context.Background(), key, value, time.Minute*2).Err()
}

// 获取验证码
func (s *store) Get(id string, clear bool) string {

	key := SystemCaptcha + id

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
