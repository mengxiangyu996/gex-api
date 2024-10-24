package dal

type Config struct {
	GomrConfig  *GomrConfig
	RedisConfig *RedisConfig
}

// 初始化数据访问层
func InitDal(config *Config) {

	// 初始化数据库
	if config.GomrConfig != nil {
		initGorm(config.GomrConfig)
	}

	// 初始化 Redis
	if config.RedisConfig != nil {
		initRedis(config.RedisConfig)
	}
}
