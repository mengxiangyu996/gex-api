package response

type Profile struct {
	Id       int    `json:"id"`
	Gender   int    `json:"gender"`
	Avatar   string `json:"avatar"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	UserId   int    `json:"userId"`
	Nickname string `json:"nickName"`
}
