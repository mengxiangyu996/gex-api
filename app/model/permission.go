package model

type Permission struct {
	Id          int `gorm:"autoIncrement"`
	Name        string
	Code        string
	Type        string
	ParentId    int
	Path        string
	Redirect    string
	Icon        string
	Component   string
	Layout      string
	KeepAlive   int
	Method      string
	Description string
	Show        int
	Enable      int `gorm:"default:1"`
	Order       int
}
