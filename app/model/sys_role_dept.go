package model

// 角色和部门关联
type SysRoleDept struct {
	RoleId int
	DeptId int
}

func (SysRoleDept) TableName() string {
	return "sys_role_dept"
}
