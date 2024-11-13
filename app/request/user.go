package request

type UserPage struct {
	Page
	Username string `query:"username" form:"username"`
	Gender   *int   `query:"gender" form:"gender"`
	Enable   *int   `query:"enable" form:"enable"`
}
