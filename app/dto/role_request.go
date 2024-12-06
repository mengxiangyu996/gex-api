package dto

type RolePageRequest struct {
	PageRequest
	Name   string `query:"name" form:"name"`
	Enable *int   `query:"enable" form:"enable"`
}

type RoleAddRequest struct {
	Code          string `json:"code"`
	Name          string `json:"name"`
	Enable        bool   `json:"enable"`
	PermissionIds []int  `json:"permissionIds"`
}

type RoleUpdateRequest struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Enable        bool   `json:"enable"`
	PermissionIds []int  `json:"permissionIds"`
}

type RoleUsersRemoveRequest struct {
	RoleId  int   `json:"roleId"`
	UserIds []int `json:"userIds"`
}

type RoleUsersAddRequest struct {
	RoleId  int   `json:"roleId"`
	UserIds []int `json:"userIds"`
}
