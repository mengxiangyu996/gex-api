package database

import (
	"breeze-api/pkg/db"
	"io/ioutil"
)

// 填充数据
func Handle() {

	sql, err := ioutil.ReadFile("database/init.sql")
	if err != nil {
		panic(err)
	}

	err = db.GormClient.Exec(string(sql)).Error
	if err != nil {
		panic(err)
	}
}