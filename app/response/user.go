package response

import (
	"isme-go/framework/datetime"
	"isme-go/types"
)

type User struct {
	Id         int               `json:"id"`
	Username   string            `json:"username"`
	Enable     types.Boolean     `json:"enable"`
	CreateTime datetime.Datetime `json:"createTime"`
	UpdateTime datetime.Datetime `json:"updateTime"`
	Password   string
}

type UserToken struct {
	Id              int      `json:"id"`
	Username        string   `json:"username"`
	RoleCodes       []string `json:"roleCodes"`
	CurrentRoleCode string   `json:"currentRoleCode"`
}

type UserPage struct {
	User
	Gender  int    `json:"gender"`
	Avatar  string `json:"avatar"`
	Address string `json:"address"`
	Email   string `json:"email"`
	Roles   []Role `json:"roles"`
}

type UserDetail struct {
	User
	Roles       []Role  `json:"roles"`
	Profile     Profile `json:"profile"`
	CurrentRole Role    `json:"currentRole"`
}
