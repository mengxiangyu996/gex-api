package response

// Option 选项响应提
type Option struct {
	Label string      `json:"label"`
	Value interface{} `json:"value"`
}

// List 分页响应体
type List struct {
	List  interface{} `json:"list"`
	Total int         `json:"total"`
}
