package response

// 响应码
var (
	SUCCESS      = 200
	ERROR        = 500
	UNAUTHORIZED = 401
)

// 响应消息
var Message = map[int]string{
	SUCCESS:      "操作成功",
	ERROR:        "操作失败",
	UNAUTHORIZED: "未授权",
}
