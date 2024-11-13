package model

type Role struct {
	Id     int `gorm:"autoIncrement"`
	Code   string
	Name   string
	Enable bool
}
