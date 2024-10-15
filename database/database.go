package database

import (
	"gex-api/pkg/dal"
	"os"
	"strings"
)

var lock = "app.lock"

func Init() {

	if _, err := os.Stat(lock); err == nil {
		return
	}

	sqlcontent, err := os.ReadFile("database/init.sql")
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
}
