package response

import "gex-api/pkg/datetime"

// 用户列表
type UserList struct {
	Id         int               `json:"id"`
	Role       int               `json:"role"`
	Username   string            `json:"username"`
	Nickname   string            `json:"nickname"`
	Gender     int               `json:"gender"`
	Email      string            `json:"email"`
	Phone      string            `json:"phone"`
	Avatar     string            `json:"avatar"`
	Status     int               `json:"status"`
	CreateTime datetime.Datetime `json:"createTime"`
}

// 用户详情
type UserDetail struct {
	Id       int    `json:"id"`
	Role     int    `json:"role"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Gender   int    `json:"gender"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar"`
	Status   int    `json:"status"`
	Password string `json:"password,omitempty"`
}
