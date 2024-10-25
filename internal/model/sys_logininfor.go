package model

import "ruoyi-go/framework/datetime"

// 系统访问记录
type SysLogininfor struct {
	InfoId        int `gorm:"autoIncrement"`
	UserName      string
	Ipaddr        string
	LoginLocation string
	Browser       string
	Os            string
	Status        string `gorm:"default:0"`
	Msg           string
	LoginTime     datetime.Datetime `gorm:"autoCreateTime"`
}

func (SysLogininfor) TableName() string {
	return "sys_logininfor"
}
