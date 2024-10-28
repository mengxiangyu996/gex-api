package logfactory

// 日志工厂
type LogFactory interface {
	NewLog() Log
}

// 登录日志工厂
type LoginLogFactory struct {
	LogFactory
}

// 初始化登录日志工厂
func (l *LoginLogFactory) NewLog() Log {
	return &LoginLog{}
}

// 操作日志工厂
type OperateLogFactory struct {
	LogFactory
}

// 初始化操作日志工厂
func (o *OperateLogFactory) NewLog() Log {
	return &OperateLog{}
}
