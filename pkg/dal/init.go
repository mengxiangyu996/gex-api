package dal

import (
	"strconv"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Config struct {
	DBConfig    *DBConfig
	RedisConfig *RedisConfig
}

type DBConfig struct {
	Dialector gorm.Dialector
	Opts      gorm.Option
}

type RedisConfig struct {
	Host     string
	Port     int
	Database int
	Password string
}

// 初始化数据访问层
func InitDal(config *Config) {

	// 初始化数据库
	if config.DBConfig != nil {
		initDB(config.DBConfig.Dialector, config.DBConfig.Opts)
	}

	// 初始化 Redis
	if config.RedisConfig != nil {
		initRedis(&redis.Options{
			Addr:     config.RedisConfig.Host + ":" + strconv.Itoa(config.RedisConfig.Port),
			Password: config.RedisConfig.Password,
			DB:       config.RedisConfig.Database,
		})
	}

	return
}
