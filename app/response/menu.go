package response

import "gex-api/pkg/datetime"

// 菜单列表
type MenuList struct {
	Id         int               `json:"id"`
	ParentId   int               `json:"parentId"`
	Name       string            `json:"name"`
	Type       int               `json:"type"`
	Sort       int               `json:"sort"`
	Path       string            `json:"path"`
	Component  string            `json:"component"`
	Icon       string            `json:"icon"`
	Hidden     int               `json:"hidden"`
	KeepAlive  int               `json:"keepAlive"`
	Redirect   string            `json:"redirect"`
	Status     int               `json:"status"`
	CreateTime datetime.Datetime `json:"createTime"`
}

// 菜单详情
type MenuDetail struct {
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

// 菜单树形列表
type MenuTree struct {
	*MenuList
	Children []*MenuTree `json:"children"`
}
