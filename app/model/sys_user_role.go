package model

// 用户和角色关联
type SysUserRole struct {
	UserId int
	RoleId int
}

func (SysUserRole) TableName() string {
	return "sys_user_role"
}
