package apiv2

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

type Captcha struct {
	Id string `json:"id"`
	Base64 string `json:"base64"`
}

type CaptchaDto struct {
	Captcha Captcha `json:"captcha"`
}

func GetCaptchaApi(c *gin.Context)  {
	store := base64Captcha.DefaultMemStore
	driver := base64Captcha.NewDriverDigit(80, 240, 6, 0.7, 80)
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, base64, err := captcha.Generate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg": "success",
			"data": CaptchaDto{
				Captcha: Captcha{
					Id: id,
					Base64: base64,
				},
			},
		})
	}
}
