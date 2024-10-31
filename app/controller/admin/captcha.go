package admin

import (
	"ruoyi-go/framework/message"
	"ruoyi-go/utils/captcha"

	"github.com/gin-gonic/gin"
)

type Captcha struct{}

// 生成验证码
func (c *Captcha) Image(ctx *gin.Context) {

	uuid, b64s := captcha.NewCaptcha().Generate()

	// 由前端拼接，this.codeUrl = "data:image/gif;base64," + res.img;
	// b64s = strings.Split(b64s, "base64,")[1]

	message.Success(ctx, map[string]interface{}{
		"uuid":           uuid,
		"img":            b64s,
		"captchaEnabled": true,
	})
}
