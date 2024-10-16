package database

import (
	"gex-api/app/internal/encrypt"
	"gex-api/app/request"
	"gex-api/app/service"
	"gex-api/pkg/dal"
	"os"
	"strings"
)

var lock = "app.lock"

func Init() {

	if _, err := os.Stat(lock); err == nil {
		return
	}

	sqlcontent, err := os.ReadFile("app/database/init.sql")
	if err != nil {
		panic(err)
	}

	sqls := strings.Split(string(sqlcontent), ";")
	for _, sql := range sqls {
		sql = strings.TrimSpace(sql)
		if sql == "" {
			continue
		}
		if err := dal.Gorm.Exec(sql).Error; err != nil {
			panic(err)
		}
	}

	file, _ := os.Create(lock)
	defer file.Close()

	// 存在超级管理员不执行生成初始超级管理员
	user := (&service.User{}).GetDetailByUsername("admin")
	if user.Id > 0 {
		return
	}

	// 生成初始超级管理员
	(&service.User{}).Create(&request.CreateUser{
		Role:     2,
		Username: "admin",
		Nickname: "超级管理员",
		Gender:   1,
		Password: encrypt.Generate("123456"),
	})
}
