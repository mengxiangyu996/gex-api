package model

import (
	"ruoyi-go/pkg/datetime"

	"gorm.io/gorm"
)

// 用户信息
type SysUser struct {
	UserId      int `gorm:"autoIncrement"`
	DeptId      int
	UserName    string
	NickName    string
	UserType    string `gorm:"default:00"`
	Email       string
	Phonenumber string
	Sex         string `gorm:"default:0"`
	Avatar      string
	Password    string
	Status      string `gorm:"default:0"`
	LoginIp     string
	LoginDate   datetime.Datetime `gorm:"default:null"`
	CreateBy    string
	CreateTime  datetime.Datetime
	UpdateBy    string
	UpdateTime  datetime.Datetime
	Remark      string
	DeleteTime  gorm.DeletedAt
}

func (SysUser) TableName() string {
	return "sys_user"
}
