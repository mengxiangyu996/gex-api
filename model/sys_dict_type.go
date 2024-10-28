package model

import "ruoyi-go/framework/datetime"

// 字典类型
type SysDictType struct {
	DictId     int `gorm:"autoIncrement"`
	DictName   string
	DictType   string
	Status     string `gorm:"default:0"`
	CreateBy   string
	CreateTime datetime.Datetime `gorm:"autoCreateTime"`
	UpdateBy   string
	UpdateTime datetime.Datetime `gorm:"autoUpdateTime"`
	Remark     string
}

func (SysDictType) TbaleName() string {
	return "sys_dict_type"
}
