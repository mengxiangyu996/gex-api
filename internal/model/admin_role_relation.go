package model

import (
	"time"

	"gorm.io/gorm"
)

// 管理员角色关系模型
type AdminRoleRelation struct {
	Id         int            `json:"id" gorm:"autoIncrement"`
	CreateTime time.Time      `json:"createTime" gorm:"autoCreateTime"`
	UpdateTime time.Time      `json:"updateTime" gorm:"autoUpdateTime"`
	DeleteTime gorm.DeletedAt `json:"deleteTime"`
	AdminId    int            `json:"adminId"`
	RoleId     int            `json:"roleId"`
}
