package response

import "isme-go/types"

type Permission struct {
	Id          int           `json:"id"`
	Name        string        `json:"name"`
	Code        string        `json:"code"`
	Type        string        `json:"type"`
	ParentId    int           `json:"parentId"`
	Path        string        `json:"path"`
	Redirect    string        `json:"redirect"`
	Icon        string        `json:"icon"`
	Component   string        `json:"component"`
	Layout      string        `json:"layout"`
	KeepAlive   types.Boolean `json:"keepAlive"`
	Method      string        `json:"method"`
	Description string        `json:"description"`
	Show        types.Boolean `json:"show"`
	Enable      types.Boolean `json:"enable"`
	Order       int           `json:"order"`
}

type PermissionTree struct {
	Permission
	Children []PermissionTree `json:"children"`
}
