package model

import (
	"time"

	"gorm.io/gorm"
)

// 管理员模型
type Admin struct {
	Id         int            `json:"id" gorm:"autoIncrement"`
	CreateTime time.Time      `json:"createTime" gorm:"autoCreateTime"`
	UpdateTime time.Time      `json:"updateTime" gorm:"autoUpdateTime"`
	DeleteTime gorm.DeletedAt `json:"deleteTime"`
	Username   string         `json:"username"`
	Nickname   string         `json:"nickname"`
	Gender     int            `json:"gender"`
	Email      string         `json:"email"`
	Phone      string         `json:"phone"`
	Password   string         `json:"password"`
	Avatar     string         `json:"avatar"`
	Status     int            `json:"status" gorm:"default:1"`
}
