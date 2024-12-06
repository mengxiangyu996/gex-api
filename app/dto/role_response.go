package dto

type RoleResponse struct {
	Id     int    `json:"id"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	Enable bool   `json:"enable"`
}

type RolePageResponse struct {
	RoleResponse
	PermissionIds []int `json:"permissionIds"`
}
