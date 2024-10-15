package request

// 创建权限
type CreatePermission struct {
	Name      string `json:"name"`
	GroupName string `json:"groupName"`
	Path      string `json:"path"`
	Method    string `json:"method"`
	Status    int    `json:"status"`
}

// 更新权限
type UpdatePermission struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	GroupName string `json:"groupName"`
	Path      string `json:"path"`
	Method    string `json:"method"`
	Status    int    `json:"status"`
}

// 权限列表
type QueryListPermission struct {
	*QueryPage
	Name      string `query:"name"`
	GroupName string `query:"groupName"`
	Path      string `query:"path"`
	Method    string `query:"method"`
	Status    int    `query:"status"`
}
