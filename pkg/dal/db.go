package dal

import (
	"time"

	"gorm.io/gorm"
)

var DBClient *gorm.DB

func initDB(dialector gorm.Dialector, opts gorm.Option) {

	var err error

	DBClient, err = gorm.Open(dialector, opts)
	if err != nil {
		panic(err)
	}

	sqlDB, err := DBClient.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	err = sqlDB.Ping()
	if err != nil {
		panic(err)
	}

	return
}
