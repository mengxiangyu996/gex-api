package model

import (
	"isme-go/framework/datetime"
)

type User struct {
	Id         int `gorm:"autoIncrement"`
	Username   string
	Password   string
	Enable     int `gorm:"default:1"`
	CreateTime datetime.Datetime
	UpdateTime datetime.Datetime
}
