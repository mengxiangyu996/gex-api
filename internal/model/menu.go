package model

import (
	"time"

	"gorm.io/gorm"
)

// 菜单模型
type Menu struct {
	Id         int            `json:"id" gorm:"autoIncrement"`
	CreateTime time.Time      `json:"createTime"`
	UpdateTime time.Time      `json:"updateTime"`
	DeleteTime gorm.DeletedAt `json:"deleteTime"`
	ParentId   int            `json:"parentId"`
	Name       string         `json:"name"`
	Type       int            `json:"type"`
	Sort       int            `json:"sort"`
	Path       string         `json:"path"`
	Component  string         `json:"component"`
	Icon       string         `json:"icon"`
	Redirect   string         `json:"redirect"`
	Status     int            `json:"status" gorm:"default:1"`
}
