package service

import (
	"ruoyi-go/framework/dal"
	"ruoyi-go/model"
	"ruoyi-go/request"
)

// 登录日志
type SysLogininfor struct{}

// 记录登录日志
func (*SysLogininfor) Insert(param *request.SysLogininforInsertRequest) error {

	data := model.SysLogininfor{
		UserName:      param.UserName,
		Ipaddr:        param.Ipaddr,
		LoginLocation: param.LoginLocation,
		Browser:       param.Browser,
		Os:            param.Os,
		Status:        param.Status,
		Msg:           param.Msg,
	}

	return dal.Gorm.Model(&model.SysLogininfor{}).Create(&data).Error
}
