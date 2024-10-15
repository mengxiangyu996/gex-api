package response

import "gex-api/pkg/datetime"

// 角色列表
type RoleList struct {
	Id         int               `json:"id"`
	Name       string            `json:"name"`
	Status     int               `json:"status"`
	CreateTime datetime.Datetime `json:"createTime"`
}

// 角色详情
type RoleDetail struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}
