package model

import (
	"time"

	"gorm.io/gorm"
)

// 角色菜单关系模型
type RoleMenuRelation struct {
	Id         int            `json:"id" gorm:"autoIncrement"`
	CreateTime time.Time      `json:"createTime" gorm:"autoCreateTime"`
	UpdateTime time.Time      `json:"updateTime" gorm:"autoUpdateTime"`
	DeleteTime gorm.DeletedAt `json:"deleteTime"`
	RoleId     int            `json:"roleId"`
	MenuId     int            `json:"menuId"`
}
