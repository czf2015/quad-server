// https://github.com/mojocn/base64Captcha#233--core-code-captchago
package api_v3

import (
	"goserver/libs/captcha"
	"goserver/libs/e"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Captcha struct {
	ID     string `json:"id"`
	Base64 string `json:"base64"`
}

func GetCaptcha(c *gin.Context) {
	code := e.SUCCESS
	if id, base64, err := captcha.Generate(); err != nil {
		code = e.ERROR_GENERATE_CAPTCHA
		c.JSON(http.StatusBadRequest, gin.H{"code": code, "message": e.GetMsg(code)})
	} else {
		c.JSON(200, gin.H{
			"code":    code,
			"message": e.GetMsg(code),
			"data": Captcha{
				ID:     id,
				Base64: base64,
			},
		})
	}
}
