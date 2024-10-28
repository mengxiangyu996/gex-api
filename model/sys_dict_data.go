package model

import "ruoyi-go/framework/datetime"

// 字典数据
type SysDictData struct {
	DictCode   int `gorm:"autoIncrement"`
	DictSort   int `gorm:"default:0"`
	DictLabel  string
	DictValue  string
	DictType   string
	CssClass   string
	ListClass  string
	IsDefault  string `gorm:"default:N"`
	Status     string `gorm:"default:0"`
	CreateBy   string
	CreateTime datetime.Datetime `gorm:"autoCreateTime"`
	UpdateBy   string
	UpdateTime datetime.Datetime `gorm:"autoUpdateTime"`
	Remark     string
}

func (SysDictData) TbaleName() string {
	return "sys_dict_data"
}
