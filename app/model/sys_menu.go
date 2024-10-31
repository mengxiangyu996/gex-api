package model

import "ruoyi-go/framework/datetime"

// 菜单权限
type SysMenu struct {
	MenuId     int64 `gorm:"autoIncrement"`
	MenuName   string
	ParentId   int64 `gorm:"default:0"`
	OrderNum   int   `gorm:"default:0"`
	Path       string
	Component  string
	Query      string
	RouteName  string
	IsFrame    int `gorm:"default:1"`
	IsCache    int `gorm:"default:0"`
	MenuType   string
	Visible    string `gorm:"default:'0'"`
	Status     string `gorm:"default:'0'"`
	Perms      string
	Icon       string `gorm:"default:'#'"`
	CreateBy   string
	CreateTime datetime.Datetime `gorm:"autoCreateTime"`
	UpdateBy   string
	UpdateTime datetime.Datetime `gorm:"autoUpdateTime"`
	Remark     string
}

func (SysMenu) TableName() string {
	return "sys_menu"
}
