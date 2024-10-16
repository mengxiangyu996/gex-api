package response

import "gex-api/pkg/datetime"

// 权限列表
type PermissionList struct {
	Id         int               `json:"id"`
	Name       string            `json:"name"`
	GroupName  string            `json:"groupName"`
	Path       string            `json:"path"`
	Method     string            `json:"method"`
	Status     int               `json:"status"`
	CreateTime datetime.Datetime `json:"createTime"`
	IsBind     bool              `json:"isBind" gorm:"-"`
}

// 权限详情
type PermissionDetail struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	GroupName string `json:"groupName"`
	Path      string `json:"path"`
	Method    string `json:"method"`
	Status    int    `json:"status"`
}
