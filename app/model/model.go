package model

import (
	"gex-api/pkg/datetime"

	"gorm.io/gorm"
)

// BaseModel 基础模型
type BaseModel struct {
	Id         int               `gorm:"autoIncrement"`
	CreateTime datetime.Datetime `gorm:"autoCreateTime"`
	UpdateTime datetime.Datetime `gorm:"autoUpdateTime"`
	DeleteTime gorm.DeletedAt
}
