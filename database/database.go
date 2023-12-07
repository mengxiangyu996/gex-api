package database

import (
	"breeze-api/helper/encrypt"
	"breeze-api/internal/model"
	"breeze-api/internal/service"
)

// 初始化数据表
func Handle() {

	(&service.User{}).CreateUser(&model.User{
		Username: "admin",
		Password: encrypt.Generate("123456"),
		IsAdmin: 1,
		Nickname: "超级管理员",
		Gender: 1,
	})

}