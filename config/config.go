package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// 项目配置
type config struct {
	// 项目相关配置
	App struct {
		Name string `yaml:"name"` // 名称
		// 开发环境配置
		Server struct {
			Port int    `yaml:"port"` // 服务器的HTTP端口
			Mode string `yaml:"mode"` // 模式
		} `yaml:"server"`
	} `yaml:"app"`

	// 数据库配置
	Mysql struct {
		Host         string `yaml:"host"`         // 地址
		Port         int    `yaml:"port"`         // 端口
		Database     string `yaml:"database"`     // 数据库名称
		Username     string `yaml:"username"`     // 用户名
		Password     string `yaml:"password"`     // 密码
		Charset      string `yaml:"charset"`      // 编码
		MaxIdleConns int    `yaml:"maxIdleConns"` // 连接池最大连接数
		MaxOpenConns int    `yaml:"maxOpenConns"` // 连接池最大打开连接数
	}

	// redis配置
	Redis struct {
		Host     string `yaml:"host"`     // 地址
		Port     int    `yaml:"port"`     // 端口
		Database int    `yaml:"database"` // 数据库索引
		Password string `yaml:"password"` // 密码
	}

	// token配置
	Token struct {
		Header     string `yaml:"header"`     // 令牌自定义标识
		Secret     string `yaml:"secret"`     // 令牌密钥
		ExpireTime int    `yaml:"expireTime"` // 令牌有效期
	}
}

var Data *config

func InitConfig() {

	file, err := os.ReadFile("application.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(file, &Data)
	if err != nil {
		panic(err)
	}
}
