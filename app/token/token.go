package token

import (
	"errors"
	"isme-go/app/dto"
	"isme-go/config"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// 用户声明
type UserClaims struct {
	UserId          int      `json:"userId"`
	Username        string   `json:"username"`
	RoleCodes       []string `json:"roleCodes"`
	CurrentRoleCode string   `json:"currentRoleCode"`
	jwt.RegisteredClaims
}

// 获取用户声明
func GetClaims(user dto.UserTokenResponse) *UserClaims {
	return &UserClaims{
		UserId:          user.Id,
		Username:        user.Username,
		RoleCodes:       user.RoleCodes,
		CurrentRoleCode: user.CurrentRoleCode,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(config.Data.Token.ExpireTime))), // 过期时间，默认30分钟
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                                                // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                                                // 生效时间
			Issuer:    "isme-go",                                                                                     // 签发人
		},
	}
}

// 生成token
func (uc *UserClaims) GenerateToken() string {

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, uc).SignedString([]byte(config.Data.Token.Secret))
	if err != nil {
		return ""
	}

	return token
}

// 解析token
func ParseToken(ctx *gin.Context) (*UserClaims, error) {

	authorization := ctx.GetHeader("Authorization")
	if authorization == "" {
		return nil, errors.New("Authorization is empty")
	}

	tokenSplit := strings.Split(authorization, " ")
	if len(tokenSplit) != 2 || tokenSplit[0] != "Bearer" {
		return nil, errors.New("authorization format error")
	}

	token, err := jwt.ParseWithClaims(tokenSplit[1], &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Data.Token.Secret), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("token格式错误")
			}
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token已过期")
			}
			if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token未生效")
			}
			return nil, errors.New("token校验失败")
		}
		return nil, err
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("token校验失败")
}
