package dto

import (
	"isme-go/framework/datetime"
)

type UserResponse struct {
	Id         int               `json:"id"`
	Username   string            `json:"username"`
	Enable     bool              `json:"enable"`
	CreateTime datetime.Datetime `json:"createTime"`
	UpdateTime datetime.Datetime `json:"updateTime"`
	Password   string
}

type UserTokenResponse struct {
	Id              int      `json:"id"`
	Username        string   `json:"username"`
	RoleCodes       []string `json:"roleCodes"`
	CurrentRoleCode string   `json:"currentRoleCode"`
}

type UserPageResponse struct {
	UserResponse
	Gender  int            `json:"gender"`
	Avatar  string         `json:"avatar"`
	Address string         `json:"address"`
	Email   string         `json:"email"`
	Roles   []RoleResponse `json:"roles"`
}

type UserDetailResponse struct {
	UserResponse
	Roles       []RoleResponse  `json:"roles"`
	Profile     ProfileResponse `json:"profile"`
	CurrentRole RoleResponse    `json:"currentRole"`
}
