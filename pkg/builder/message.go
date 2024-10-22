package builder

// 状态码
const (
	SUCCESS_STATUS = 10200
	FAIL_STATUS    = 10500
)

type Message struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func json(code int, message string, data interface{}) *Message {
	return &Message{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
