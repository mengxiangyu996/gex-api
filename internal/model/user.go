package model

import (
	"time"

	"gorm.io/gorm"
)

// 用户模型
type User struct {
	Id         int            `json:"id" gorm:"autoIncrement"`
	CreateTime time.Time      `json:"createTime"`
	UpdateTime time.Time      `json:"updateTime"`
	DeleteTime gorm.DeletedAt `json:"deleteTime"`
	IsAdmin    int            `json:"isAdmin" gorm:"default:0"`
	Username   string         `json:"username"`
	Nickname   string         `json:"nickname"`
	Gender     int            `json:"gender"`
	Email      string         `json:"email"`
	Phone      string         `json:"phone"`
	Password   string         `json:"password"`
	Avatar     string         `json:"avatar"`
	WxOpenId   string         `json:"wxOpenId"`
	WxUnionId  string         `json:"wxUnionId"`
	Status     int            `json:"status" gorm:"default:1"`
}
