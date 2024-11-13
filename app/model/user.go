package model

import (
	"isme-go/framework/datetime"
)

type User struct {
	Id         int `gorm:"autoIncrement"`
	Username   string
	Password   string
	Enable     int
	CreateTime datetime.Datetime `gorm:"autoCreateTime"`
	UpdateTime datetime.Datetime `gorm:"autoUpdateTime"`
}
