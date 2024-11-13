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
	KeepAlive   bool
	Method      string
	Description string
	Show        bool
	Enable      bool
	Order       int
}
