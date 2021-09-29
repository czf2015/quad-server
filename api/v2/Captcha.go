// https://github.com/mojocn/base64Captcha#233--core-code-captchago
package api_v2

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"goserver/libs/captcha"
)

type Captcha struct {
	Id string `json:"id"`
	Base64 string `json:"base64"`
}

type CaptchaResponse struct {
	Captcha Captcha `json:"captcha"`
}

func GetCaptchaApi(c *gin.Context)  {
	id, base64, err := captcha.Generate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg": "success",
			"data": CaptchaResponse{
				Captcha: Captcha{
					Id: id,
					Base64: base64,
				},
			},
		})
	}
}
