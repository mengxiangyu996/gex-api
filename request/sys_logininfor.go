package request

// 创建登录日志
type SysLogininforInsertRequest struct {
	UserName      string
	Ipaddr        string
	LoginLocation string
	Browser       string
	Os            string
	Status        string
	Msg           string
}
