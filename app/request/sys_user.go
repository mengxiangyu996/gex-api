package request

type SysUserLogin struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	Uuid     string `json:"uuid"`
	Code     string `json:"code"`
}
