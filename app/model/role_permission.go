package model

// 用户权限
type RolePermission struct {
	BaseModel
	RoleId       int
	PermissionId int
}
