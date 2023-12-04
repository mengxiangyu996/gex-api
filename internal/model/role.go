package model

import (
	"time"

	"gorm.io/gorm"
)

// 角色模型
type Role struct {
	Id         int            `json:"id" gorm:"autoIncrement"`
	CreateTime time.Time      `json:"createTime"`
	UpdateTime time.Time      `json:"updateTime"`
	DeleteTime gorm.DeletedAt `json:"deleteTime"`
	Name       string         `json:"name"`
	Status     int            `json:"status" gorm:"default:1"`
}
