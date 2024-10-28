package captcha

import (
	"ruoyi-go/config"

	"github.com/mojocn/base64Captcha"
)

// 验证码
type Captcha struct {
	captcha *base64Captcha.Captcha
}

// 创建验证码对象
func NewCaptcha() *Captcha {

	store := NewStore()

	var driver base64Captcha.Driver

	switch config.Data.RuoYi.CaptchaType {
	case "math":
		driver = base64Captcha.NewDriverMath(60, 160, 200, base64Captcha.OptionShowHollowLine, nil, nil, nil)
	case "char":
		source := "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		driver = base64Captcha.NewDriverString(60, 160, 200, base64Captcha.OptionShowHollowLine, 4, source, nil, nil, nil)
	default:
		driver = base64Captcha.DefaultDriverDigit
	}

	return &Captcha{
		captcha: base64Captcha.NewCaptcha(driver, store),
	}
}

// 生成验证码
// uuid, base64, answer
func (c *Captcha) Generate() (string, string) {

	id, b64s, _, err := c.captcha.Generate()
	if err != nil {
		return "", ""
	}

	return id, b64s
}

// 验证验证码
func (c *Captcha) Verify(id, answer string) bool {
	return c.captcha.Verify(id, answer, true)
}
