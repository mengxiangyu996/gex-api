package model

import (
	"time"

	"gorm.io/gorm"
)

// 权限模型
type Permission struct {
	Id         int            `json:"id" gorm:"autoIncrement"`
	CreateTime time.Time      `json:"createTime"`
	UpdateTime time.Time      `json:"updateTime"`
	DeleteTime gorm.DeletedAt `json:"deleteTime"`
	Name       string         `json:"name"`
	Path       string         `json:"path"`
	Method     string         `json:"method"`
	Status     int            `json:"status" gorm:"default:1"`
}
