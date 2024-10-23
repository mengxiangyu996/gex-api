package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// 项目配置
type config struct {
	// 项目相关配置
	RuoYi struct {
		Name           string `yaml:"name"`           // 名称
		Version        string `yaml:"version"`        // 版本
		CopyrightYear  int    `yaml:"copyrightYear"`  // 版权年份
		Profile        string `yaml:"profile"`        // 文件路径
		AddressEnabled bool   `yaml:"addressEnabled"` // 获取ip地址开关
		CaptchaType    string `yaml:"captchaType"`    // 验证码类型
	}

	// 开发环境配置
	Server struct {
		Port int    `yaml:"port"` // 服务器的HTTP端口
		Mode string `yaml:"mode"` // 模式
	}

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

	// 用户配置
	User struct {
		Password struct {
			MaxRetryCount int `yaml:"maxRetryCount"` // 密码最大错误次数
			LockTime      int `yaml:"lockTime"`      // 密码锁定时间
		} `yaml:"password"`
	}

	// swagger配置
	Swagger struct {
		Enabled     bool   `yaml:"enabled"`     // 是否开启swagger
		PathMapping string `yaml:"pathMapping"` // 请求前缀
	} `yaml:"swagger"`
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
