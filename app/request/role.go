package request

// 创建角色
type CreateRole struct {
	Name   string `json:"name"`
	Status int    `json:"status"`
}

// 更新角色
type UpdateRole struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

// 角色列表
type QueryListRole struct {
	*QueryPage
	Name   string `query:"name"`
	Status int    `query:"status"`
}
