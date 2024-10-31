package response

import "ruoyi-go/framework/datetime"

type SysUserDetail struct {
	UserId      int               `json:"userId"`
	DeptId      int               `json:"deptId"`
	UserName    string            `json:"userName"`
	NickName    string            `json:"nickName"`
	UserType    string            `json:"userType"`
	Email       string            `json:"email"`
	Phonenumber string            `json:"phonenumber"`
	Sex         string            `json:"sex"`
	Avatar      string            `json:"avatar"`
	Password    string            `json:"-"`
	Status      string            `json:"status"`
	LoginIp     string            `json:"loginIp"`
	LoginDate   datetime.Datetime `json:"loginDate"`
	CreateBy    string            `json:"createBy"`
	CreateTime  datetime.Datetime
	UpdateBy    string            `json:"updateBy"`
	UpdateTime  datetime.Datetime `json:"updateTime"`
	Remark      string            `json:"remark"`
}
