package dal

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type DBConfig struct {
	GormConfig  *GormConfig
	RedisConfig *RedisConfig
}

var (
	Gorm  *gorm.DB
	Redis *redis.Client
)

func Init(config *DBConfig) {

	// 初始化Gorm
	if config.GormConfig != nil {
		initGorm(config.GormConfig)
	}

	// 初始化Redis
	if config.RedisConfig != nil {
		initRedis(&redis.Options{
			Addr:     config.RedisConfig.Address,
			Password: config.RedisConfig.Password,
			DB:       config.RedisConfig.Database,
		})
	}
}
