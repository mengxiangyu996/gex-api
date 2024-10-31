package model

// 用户和岗位关联
type SysUserPost struct {
	UserId int
	PostId int
}

func (SysUserPost) TableName() string {
	return "sys_user_post"
}
