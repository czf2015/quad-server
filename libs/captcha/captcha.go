package captcha

import (
	"github.com/mojocn/base64Captcha"
)

var captcha *base64Captcha.Captcha

func init() {
	store := base64Captcha.DefaultMemStore
	driver := base64Captcha.NewDriverDigit(80, 240, 6, 0.7, 80)
	captcha = base64Captcha.NewCaptcha(driver, store)
}

func Generate() (id, b64s string, err error) {
	return captcha.Generate()
}

func Verify(id, answer string, clear bool) bool {
	return captcha.Verify(id, answer, clear)
}