package model

import (
	"time"

	"gorm.io/gorm"
)

// 角色权限关系模型
type RolePermissionRelation struct {
	Id           int            `json:"id" gorm:"autoIncrement"`
	CreateTime   time.Time      `json:"createTime"`
	UpdateTime   time.Time      `json:"updateTime"`
	DeleteTime   gorm.DeletedAt `json:"deleteTime"`
	RoleId       int            `json:"roleId"`
	PermissionId int            `json:"permissionId"`
}
