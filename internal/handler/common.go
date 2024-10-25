package handler

import (
	"ruoyi-go/framework/response"
	"ruoyi-go/utils/captcha"
	"strings"

	"github.com/gin-gonic/gin"
)

type Common struct{}

// 生成验证码
func (c *Common) CaptchaImage(ctx *gin.Context) {

	uuid, b64s := captcha.NewCaptcha().Generate()

	// 由前端拼接，this.codeUrl = "data:image/gif;base64," + res.img;
	b64s = strings.Split(b64s, "base64,")[1]

	response.Success(ctx, map[string]interface{}{
		"uuid":           uuid,
		"img":            b64s,
		"captchaEnabled": true,
	})
}
