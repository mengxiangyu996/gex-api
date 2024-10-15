package controller

import (
	"gex-api/app/internal/encrypt"
	"gex-api/app/internal/jwt"
	"gex-api/app/internal/utils"
	"gex-api/app/request"
	"gex-api/app/response"
	"gex-api/app/service"
	"gex-api/pkg/builder"
	"time"
)

// 用户
type User struct{}

// 创建用户
func (*User) Create(ctx *builder.Context) error {

	var param request.CreateUser

	if err := ctx.BindX(&param); err != nil {
		return ctx.Fail(err.Error())
	}

	if param.Username == "" {
		return ctx.Fail("参数错误")
	}

	if user := (&service.User{}).GetDetailByUsername(param.Username); user.Id > 0 {
		return ctx.Fail("用户已存在")
	}

	param.Password = encrypt.Generate(param.Password)

	if uid := (&service.User{}).Create(&param); uid <= 0 {
		return ctx.Fail()
	}

	return ctx.Success()
}

// 更新用户
func (*User) Update(ctx *builder.Context) error {

	var param request.UpdateUser

	if err := ctx.BindX(&param); err != nil {
		return ctx.Fail(err.Error())
	}

	if user := (&service.User{}).GetDetailByUsername(param.Username); user.Id > 0 && user.Id != param.Id {
		return ctx.Fail("用户已存在")
	}

	if uid := (&service.User{}).Update(&param); uid <= 0 {
		return ctx.Fail()
	}

	return ctx.Success()
}

// 删除用户
func (*User) Delete(ctx *builder.Context) error {

	var param request.QueryId

	if err := ctx.BindX(&param); err != nil {
		return ctx.Fail(err.Error())
	}

	if err := (&service.User{}).DeleteById(param.Id); err != nil {
		return ctx.Fail(err.Error())
	}

	return ctx.Success()
}

// 用户列表
func (*User) List(ctx *builder.Context) error {

	var param request.QueryListUser

	if err := ctx.BindX(&param); err != nil {
		return ctx.Fail(err.Error())
	}

	list, count := (&service.User{}).GetList(&param)

	return ctx.Success(&response.List{
		List:  list,
		Total: count,
	})
}

// 用户详情
func (*User) Detail(ctx *builder.Context) error {

	var param request.QueryId

	if err := ctx.BindX(&param); err != nil {
		return ctx.Fail(err.Error())
	}

	detail := (&service.User{}).GetDetailById(param.Id)

	return ctx.Success(detail)
}

// 后台用户登录
func (*User) Login(ctx *builder.Context) error {

	var param request.AdminUserLogin

	if err := ctx.BindX(&param); err != nil {
		return ctx.Fail(err.Error())
	}

	if param.Username == "" || param.Password == "" {
		return ctx.Fail("参数错误")
	}

	user := (&service.User{}).GetDetailByUsername(param.Username)
	if user.Id <= 0 || user.Role != 2 {
		return ctx.Fail("用户不存在")
	}

	if !encrypt.Compare(user.Password, param.Password) {
		return ctx.Fail("密码错误")
	}

	token := jwt.Generate(&jwt.Payload{
		Id:     user.Id,
		Expire: time.Now().AddDate(0, 0, 7),
	})

	return ctx.Success(token)
}

// 后台用户重置密码
func (*User) ResetPassword(ctx *builder.Context) error {

	id, _ := utils.GetTokenPayload(ctx.GetHeader("Token"))

	var param request.AdminResetPassword

	if err := ctx.BindX(&param); err != nil {
		return ctx.Fail(err.Error())
	}

	if param.Password == "" {
		return ctx.Fail("密码不能为空")
	}

	if uid := (&service.User{}).Update(&request.UpdateUser{
		Id:       id,
		Password: encrypt.Generate(param.Password),
	}); uid <= 0 {
		return ctx.Fail()
	}

	return ctx.Success()
}

// 后台用户绑定角色
func (*User) BindRole(ctx *builder.Context) error {

	var param request.AdminBindRole

	if err := ctx.BindX(&param); err != nil {
		return ctx.Fail(err.Error())
	}

	if param.UserId <= 0 {
		return ctx.Fail("参数错误")
	}

	if user := (&service.User{}).GetDetailById(param.UserId); user.Id <= 0 || user.Role != 2 {
		return ctx.Fail("用户不是管理员")
	}

	if err := (&service.User{}).BindRole(param.UserId, param.RoleIds); err != nil {
		return ctx.Fail(err.Error())
	}

	return ctx.Success()
}
