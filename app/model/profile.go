package model

type Profile struct {
	Id       int `gorm:"autoIncrement"`
	Gender   int
	Avatar   string
	Address  string
	Email    string
	UserId   int
	NickName string
}
