package model

// 用户模型
type User struct {
	BaseModel
	Role     int `gorm:"default:1"`
	Username string
	Nickname string
	Gender   int
	Email    string
	Phone    string
	Password string
	Avatar   string
	Status   int `gorm:"default:1"`
}
