package dal

import (
	"time"

	"gorm.io/gorm"
)

type GormConfig struct {
	Dialector gorm.Dialector
	Opts      gorm.Option
}

func initGorm(config *GormConfig) {

	var err error

	Gorm, err = gorm.Open(config.Dialector, config.Opts)
	if err != nil {
		panic(err)
	}

	sqlDB, err := Gorm.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err = sqlDB.Ping(); err != nil {
		panic(err)
	}
}
