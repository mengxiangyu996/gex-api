package request

// 创建菜单
type CreateMenu struct {
	ParentId  int    `json:"parentId"`
	Name      string `json:"name"`
	Type      int    `json:"type"`
	Sort      int    `json:"sort"`
	Path      string `json:"path"`
	Component string `json:"component"`
	Icon      string `json:"icon"`
	Hidden    int    `json:"hidden"`
	KeepAlive int    `json:"keepAlive"`
	Redirect  string `json:"redirect"`
	Status    int    `json:"status"`
}

// 更新菜单
type UpdateMenu struct {
	Id        int    `json:"id"`
	ParentId  int    `json:"parentId"`
	Name      string `json:"name"`
	Type      int    `json:"type"`
	Sort      int    `json:"sort"`
	Path      string `json:"path"`
	Component string `json:"component"`
	Icon      string `json:"icon"`
	Hidden    int    `json:"hidden"`
	KeepAlive int    `json:"keepAlive"`
	Redirect  string `json:"redirect"`
	Status    int    `json:"status"`
}

// 菜单列表
type QueryListMenu struct {
	QueryPage
	Name      string `query:"name"`
	Path      string `query:"path"`
	Component string `query:"component"`
}
