package request

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Captcha  string `json:"captcha"`
}

type Password struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}
