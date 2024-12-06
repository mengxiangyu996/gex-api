package dto

// 分页请求
type PageRequest struct {
	PageNo   int `query:"pageNo" form:"pageNo"`
	PageSize int `query:"pageSize" form:"pageSize"`
}
