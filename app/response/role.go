package response

type Role struct {
	Id     int    `json:"id"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	Enable bool   `json:"enable"`
}

type RolePage struct {
	Role
	PermissionIds []int `json:"permissionIds"`
}
