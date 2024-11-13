package request

// 分页请求
type Page struct {
	PageNo   int `query:"pageNo" form:"pageNo"`
	PageSize int `query:"pageSize" form:"pageSize"`
}
