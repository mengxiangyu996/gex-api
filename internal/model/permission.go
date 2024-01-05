package model

import (
	"time"

	"gorm.io/gorm"
)

// 权限模型
type Permission struct {
	Id         int            `json:"id" gorm:"autoIncrement"`
	CreateTime time.Time      `json:"createTime" gorm:"autoCreateTime"`
	UpdateTime time.Time      `json:"updateTime" gorm:"autoUpdateTime"`
	DeleteTime gorm.DeletedAt `json:"deleteTime"`
	Name       string         `json:"name"`
	GroupName  string         `json:"groupName"`
	Path       string         `json:"path"`
	Method     string         `json:"method"`
	Status     int            `json:"status" grom:"default:1"`
}
