package model

import (
	"ruoyi-go/pkg/datetime"

	"gorm.io/gorm"
)

// 部门
type SysDept struct {
	DeptId     int `gorm:"autoIncrement"`
	ParentId   int `gorm:"default:0"`
	Ancestors  string
	DeptName   string
	OrderNum   int `gorm:"default:0"`
	Leader     string
	Phone      string
	Email      string
	Status     string `gorm:"default:0"`
	CreateBy   string
	CreateTime datetime.Datetime `gorm:"autoCreateTime"`
	UpdateBy   string
	UpdateTime datetime.Datetime `gorm:"autoUpdateTime"`
	DeleteTime gorm.DeletedAt
}

func (SysDept) TbaleName() string {
	return "sys_dept"
}
