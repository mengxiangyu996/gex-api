package logfactory

import (
	"errors"
	"ruoyi-go/request"
	"ruoyi-go/service"
)

// 日志接口
type Log interface {
	Record() error
}

// 登录日志
type LoginLog struct {
	SysLoginLog *request.SysLogininforInsertRequest
}

// 记录登录日志
func (l *LoginLog) Record() error {

	if l.SysLoginLog == nil {
		return errors.New("登录日志为空")
	}

	return (&service.SysLogininfor{}).Insert(l.SysLoginLog)
}

// 操作日志
type OperateLog struct {
	SysOperLog *request.SysOperLogInsertRequest
}

func (l *OperateLog) Record() error {

	if l.SysOperLog == nil {
		return errors.New("操作日志为空")
	}

	return (&service.SysOperLog{}).Insert(l.SysOperLog)
}
