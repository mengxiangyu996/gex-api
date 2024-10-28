package service

import (
	"ruoyi-go/framework/dal"
	"ruoyi-go/model"
	"ruoyi-go/request"
)

// 操作日志
type SysOperLog struct{}

// 记录操作日志
func (*SysOperLog) Insert(param *request.SysOperLogInsertRequest) error {

	data := model.SysOperLog{
		Title:         param.Title,
		BusinessType:  param.BusinessType,
		Method:        param.Method,
		RequestMethod: param.RequestMethod,
		OperatorType:  param.OperatorType,
		OperName:      param.OperName,
		DeptName:      param.DeptName,
		OperUrl:       param.OperUrl,
		OperIp:        param.OperIp,
		OperLocation:  param.OperLocation,
		OperParam:     param.OperParam,
		JsonResult:    param.JsonResult,
		Status:        param.Status,
		ErrorMsg:      param.ErrorMsg,
		CostTime:      param.CostTime,
	}

	return dal.Gorm.Model(&model.SysOperLog{}).Create(&data).Error
}
