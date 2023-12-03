package model

import (
	"time"

	"gorm.io/gorm"
)

// 权限模型
type Permission struct {
	Id         int            `json:"id" gorm:"autoIncrement"`
	ParentId   int            `json:"parentId"`
	Name       string         `json:"name"`
	Type       int            `json:"type"`
	Path       string         `json:"path"`
	Status     int            `json:"status" gorm:"default:1"`
	CreateTime time.Time      `json:"createTime"`
	UpdateTime time.Time      `json:"updateTime"`
	DeleteTime gorm.DeletedAt `json:"deleteTime"`
}
