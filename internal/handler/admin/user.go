package admin

import (
	"breeze-api/helper/encrypt"
	"breeze-api/helper/jwt"
	"breeze-api/internal/model"
	"breeze-api/internal/service"
	"breeze-api/pkg/response"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type User struct{}

// 登录
func (t *User) Login(ctx *fiber.Ctx) error {

	type request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var req request

	ctx.BodyParser(&req)

	if req.Username == "" || req.Password == "" {
		return response.Error(ctx, "参数错误")
	}

	user := (&service.User{}).GetUserByUsername(req.Username)

	if user.Id <= 0 {
		return response.Error(ctx, "账号不存在")
	}

	if user.Status != 1 {
		return response.Error(ctx, "账号异常")
	}

	if !encrypt.Compare(user.Password, req.Password) {
		return response.Error(ctx, "密码错误")
	}

	token := jwt.Generate(&jwt.Payload{
		Id:     user.Id,
		Expire: time.Now().AddDate(0, 1, 0),
	})

	return response.Success(ctx, "ok", map[string]interface{}{
		"token": token,
	})
}

// 添加
func (t *User) CreateUser(ctx *fiber.Ctx) error {

	type request struct {
		IsAdmin  int    `json:"isAdmin"`
		Username string `json:"username"`
		Nickname string `json:"nickname"`
		Gender   int    `json:"gender"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Password string `json:"password"`
		Avatar   string `json:"avatar"`
		Status   int    `json:"status"`
	}

	var req request

	ctx.BodyParser(&req)

	if req.Username == "" || req.Password == "" {
		return response.Error(ctx, "参数错误")
	}

	user := (&service.User{}).GetUserByUsername(req.Username)
	if user.Id > 0 {
		return response.Error(ctx, "账号已存在")
	}

	err := (&service.User{}).CreateUser(&model.User{
		IsAdmin:  req.IsAdmin,
		Username: req.Username,
		Nickname: req.Nickname,
		Gender:   req.Gender,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: encrypt.Generate(req.Password),
		Avatar:   req.Avatar,
		Status:   req.Status,
	})

	if err != nil {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 更新
func (t *User) UpdateUser(ctx *fiber.Ctx) error {

	type request struct {
		Id       int    `json:"id"`
		IsAdmin  int    `json:"isAdmin"`
		Nickname string `json:"nickname"`
		Gender   int    `json:"gender"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Avatar   string `json:"avatar"`
		Status   int    `json:"status"`
	}

	var req request

	ctx.BodyParser(&req)

	if req.Id <= 0 {
		return response.Error(ctx, "参数错误")
	}

	err := (&service.User{}).UpadteUser(&model.User{
		Id:       req.Id,
		IsAdmin:  req.IsAdmin,
		Nickname: req.Nickname,
		Gender:   req.Gender,
		Email:    req.Email,
		Phone:    req.Phone,
		Avatar:   req.Avatar,
		Status:   req.Status,
	})

	if err != nil {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 删除
func (t *User) DeleteUser(ctx *fiber.Ctx) error {

	type request struct {
		Id int `json:"id"`
	}

	var req request

	ctx.BodyParser(&req)

	if req.Id <= 0 {
		return response.Error(ctx, "参数错误")
	}

	if req.Id == 1 {
		return response.Error(ctx, "超级管理员无法删除")
	}

	err := (&service.User{}).DeleteUser(req.Id)

	if err != nil {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 列表
func (t *User) GetUserPage(ctx *fiber.Ctx) error {

	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	size, _ := strconv.Atoi(ctx.Query("size", "10"))

	list, total := (&service.User{}).GetUserListByPage(page, size, -1)

	return response.Success(ctx, "成功", map[string]interface{}{
		"list":  list,
		"total": total,
	})
}

// 详情
func (t *User) GetUser(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Query("id"))

	user := (&service.User{}).GetUserById(id)

	return response.Success(ctx, "成功", user)
}

// 修改密码
func (t *User) UpdatePassword(ctx *fiber.Ctx) error {

	type request struct {
		Password string `json:"password"`
	}

	var req request

	ctx.BodyParser(&req)

	if strings.TrimSpace(req.Password) == "" {
		return response.Error(ctx, "参数错误")
	}

	userId := jwt.Parse(ctx.Get("Token")).Id

	err := (&service.User{}).UpadteUser(&model.User{
		Id:       userId,
		Password: encrypt.Generate(req.Password),
	})

	if err != nil {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}
