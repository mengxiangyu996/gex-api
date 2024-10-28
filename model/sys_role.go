package model

import (
	"ruoyi-go/framework/datetime"

	"gorm.io/gorm"
)

// 角色信息
type SysRole struct {
	RoleId            int `gorm:"autoIncrement"`
	RoleName          string
	RoleKey           string
	RoleSort          int    `gorm:"default:0"`
	DataScope         string `gorm:"default:1"`
	MenuCheckStrictly int    `gorm:"default:1"`
	DeptCheckStrictly int    `gorm:"default:1"`
	Status            string `gorm:"default:0"`
	CreateBy          string
	CreateTime        datetime.Datetime `gorm:"autoCreateTime"`
	UpdateBy          string
	UpdateTime        datetime.Datetime `gorm:"autoUpdateTime"`
	Remark            string
	DeleteTime        gorm.DeletedAt
}

func (SysRole) TableName() string {
	return "sys_role"
}
