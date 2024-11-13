package request

type RolePage struct {
	Page
	Name   string `query:"name" form:"name"`
	Enable *int   `query:"enable" form:"enable"`
}

type RoleAdd struct {
	Code          string `json:"code"`
	Name          string `json:"name"`
	Enable        bool   `json:"enable"`
	PermissionIds []int  `json:"permissionIds"`
}

type RoleUpdate struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Enable        bool   `json:"enable"`
	PermissionIds []int  `json:"permissionIds"`
}

type RoleUsersRemove struct {
	RoleId  int   `json:"roleId"`
	UserIds []int `json:"userIds"`
}

type RoleUsersAdd struct {
	RoleId  int   `json:"roleId"`
	UserIds []int `json:"userIds"`
}
