package request

// 主键id请求
type QueryId struct {
	Id int `query:"id"`
}

// 分页请求
type QueryPage struct {
	IsPaging bool `default:"true"`
	Page     int  `query:"page" default:"1"`
	Size     int  `query:"size" default:"10"`
}
