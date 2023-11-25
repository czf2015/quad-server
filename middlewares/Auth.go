package middlewares

import (
	"goserver/libs/captcha"
	"goserver/libs/e"
	"goserver/libs/utils"
	models "goserver/models/v3"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	User   models.User
	Status int
}

func CheckAuth(username, password, captchaID, captchaCode string) (auth Auth) {
	auth.Status = e.ERROR_AUTH
	if captcha.Verify(captchaID, captchaCode, true) {
		var user models.User
		db.Where(models.User{Name: username, Password: utils.EncryptPassword(password)}).First(&user)
		if user.ID > 0 {
			auth.User = user
			if user.Activated {
				auth.Status = e.SUCCESS
			} else {
				auth.Status = e.ERROR_AUTH_INACTIVE
			}
		}
	}
	return
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 在这里进行身份验证逻辑，例如检查用户的登录状态、会话等
		// 如果验证失败，可以返回适当的错误响应，并终止请求链
		// 如果验证成功，将用户信息保存到上下文中，供后续处理函数使用
		user := models.User{Name: "Demo User", Email: "demo@example.com"}
		c.Set("user", user)
		c.Next()
	}
}
