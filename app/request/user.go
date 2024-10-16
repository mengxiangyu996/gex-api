package request

// 创建用户
type CreateUser struct {
	Role     int    `json:"role"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Gender   int    `json:"gender"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	Status   int    `json:"status"`
}

// 更新用户
type UpdateUser struct {
	Id       int    `json:"id"`
	Role     int    `json:"role"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Gender   int    `json:"gender"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	Status   int    `json:"status"`
}

// 用户列表
type QueryListUser struct {
	QueryPage
	Role     int    `query:"role"`
	Username string `query:"username"`
	Nickname string `query:"nickname"`
	Phone    string `query:"phone"`
	Email    string `query:"email"`
	Status   int    `query:"status"`
}

// 后台用户登录
type AdminUserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 后台用户重置密码
type AdminResetPassword struct {
	Password string `json:"password"`
}

// 后台用户绑定角色
type AdminBindRole struct {
	UserId  int   `json:"userId"`
	RoleIds []int `json:"roleIds"`
}
