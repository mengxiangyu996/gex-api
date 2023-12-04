package model

import (
	"time"

	"gorm.io/gorm"
)

// 角色菜单关系模型
type RoleMenuRelation struct {
	Id         int            `json:"id" gorm:"autoIncrement"`
	CreateTime time.Time      `json:"createTime"`
	UpdateTime time.Time      `json:"updateTime"`
	DeleteTime gorm.DeletedAt `json:"deleteTime"`
	RoleId     int            `json:"roleId"`
	MenuId     int            `json:"menuId"`
}
