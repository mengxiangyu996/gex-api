package model

import "ruoyi-go/framework/datetime"

// 参数配置
type SysConfig struct {
	ConfigId    int `gorm:"autoIncrement"`
	ConfigName  string
	ConfigKey   string
	ConfigValue string
	ConfigType  string `gorm:"default:N"`
	CreateBy    string
	CreateTime  datetime.Datetime `gorm:"autoCreateTime"`
	UpdateBy    string
	UpdateTime  datetime.Datetime `gorm:"autoUpdateTime"`
	Remark      string
}

func (SysConfig) TableName() string {
	return "sys_config"
}
