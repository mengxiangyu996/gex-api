package model

import "ruoyi-go/framework/datetime"

// 通知公告
type SysNotice struct {
	NoticeId      int `gorm:"autoIncrement"`
	NoticeTitle   string
	NoticeType    string
	NoticeContent string
	Status        string `gorm:"default:0"`
	CreateBy      string
	CreateTime    datetime.Datetime `gorm:"autoCreateTime"`
	UpdateBy      string
	UpdateTime    datetime.Datetime `gorm:"autoUpdateTime"`
	Remark        string
}

func (SysNotice) TableName() string {
	return "sys_notice"
}
