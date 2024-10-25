package response

// 响应码
const (
	SUCCESS      = 200
	ERROR        = 500
	UNAUTHORIZED = 401
)

// 响应消息
var message = map[int]string{
	SUCCESS:      "操作成功",
	ERROR:        "操作失败",
	UNAUTHORIZED: "未授权",
}

func getMessage(code int) string {

	if msg, ok := message[code]; ok {
		return msg
	}

	return "ok"
}
