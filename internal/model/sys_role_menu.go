package model

// 角色和菜单关联
type SysRoleMenu struct {
	RoleId int
	MenuId int
}

func (SysRoleMenu) TableName() string {
	return "sys_role_menu"
}
