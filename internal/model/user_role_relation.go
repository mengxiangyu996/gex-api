package model

import (
	"time"

	"gorm.io/gorm"
)

// 用户角色关系模型
type UserRoleRelation struct {
	Id         int            `json:"id" gorm:"autoIncrement"`
	CreateTime time.Time      `json:"createTime"`
	UpdateTime time.Time      `json:"updateTime"`
	DeleteTime gorm.DeletedAt `json:"deleteTime"`
	UserId     int            `json:"userId"`
	RoleId     int            `json:"roleId"`
}
