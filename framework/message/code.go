package message

const (
	SUCCESS        = 0
	ERROR_CODE     = 10500
	ERROR_CODE_401 = 401
)

var message = map[int]string{
	SUCCESS:        "OK",
	ERROR_CODE:     "FAILED",
	ERROR_CODE_401: "登录已过期",
}

func getMessage(code int) string {

	if msg, ok := message[code]; ok {
		return msg
	}

	return "ok"
}