package model

import "ruoyi-go/framework/datetime"

// 操作日志
type SysOperLog struct {
	OperId        int `gorm:"autoIncrement"`
	Title         string
	BusinessType  string `gorm:"default:0"`
	Method        string
	RequestMethod string
	OperatorType  string `gorm:"default:0"`
	OperName      string
	DeptName      string
	OperUrl       string
	OperIp        string
	OperLocation  string
	OperParam     string
	JsonResult    string
	Status        string `gorm:"default:0"`
	ErrorMsg      string
	OperTime      datetime.Datetime `gorm:"autoCreateTime"`
	CostTime      int
}

func (SysOperLog) TableName() string {
	return "sys_oper_log"
}
