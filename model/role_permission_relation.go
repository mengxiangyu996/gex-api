package model

import (
	"time"

	"gorm.io/gorm"
)

// 角色权限关联模型
type RolePermissionRelation struct {
	Id           int            `json:"id" gorm:"autoIncrement"`
	RoleId       int            `json:"roleId"`
	PermissionId int            `json:"permissionId"`
	CreateTime   time.Time      `json:"createTime"`
	UpdateTime   time.Time      `json:"updateTime"`
	DeleteTime   gorm.DeletedAt `json:"deleteTime"`
}
