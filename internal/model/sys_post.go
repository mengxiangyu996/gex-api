package model

import "ruoyi-go/framework/datetime"

// 岗位信息
type SysPost struct {
	PostId     int `gorm:"autoIncrement"`
	PostCode   string
	PostName   string
	PostSort   int    `gorm:"default:0"`
	Status     string `gorm:"default:0"`
	CreateBy   string
	CreateTime datetime.Datetime `gorm:"autoCreateTime"`
	UpdateBy   string
	UpdateTime datetime.Datetime `gorm:"autoUpdateTime"`
	Remark     string
}

func (SysPost) TableName() string {
	return "sys_post"
}
