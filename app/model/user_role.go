package model

// 用户权限
type UserRole struct {
	BaseModel
	UserId int
	RoleId int
}
