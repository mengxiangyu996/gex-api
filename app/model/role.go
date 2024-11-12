package model

type Role struct {
	Id     int `gorm:"autoIncrement"`
	Code   string
	Name   string
	Enable int `gorm:"default:1"`
}
