package dto

type PermissionResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Type        string `json:"type"`
	ParentId    int    `json:"parentId"`
	Path        string `json:"path"`
	Redirect    string `json:"redirect"`
	Icon        string `json:"icon"`
	Component   string `json:"component"`
	Layout      string `json:"layout"`
	KeepAlive   bool   `json:"keepAlive"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Show        bool   `json:"show"`
	Enable      bool   `json:"enable"`
	Order       int    `json:"order"`
}

type PermissionTreeResponse struct {
	PermissionResponse
	Children []PermissionTreeResponse `json:"children"`
}
