package middlewares

import (
	"goserver/libs/e"
	"goserver/libs/gorm"
	"goserver/libs/utils"
	"goserver/libs/captcha"
	models "goserver/models/v2"
)

type Auth struct {
	User models.User
	Activation models.Activation
	Status int
}

func CheckAuth(userName, password, captchaID, captchaCode string) (auth Auth) {
	auth.Status = e.ERROR_AUTH
	if (captcha.Verify(captchaID, captchaCode, true)) {
		var user models.User
		gorm.GetDB().Where(models.User{Name: userName, Password: utils.EncryptPassword(password)}).First(&user)
		if len(user.ID) > 0 {
				var activation models.Activation
				gorm.GetDB().Where(models.Activation{UserId: user.ID}).Where("completed_at IS NOT NULL").First(&activation)
				auth.User = user
				auth.Activation = activation
				if len(activation.CompletedAt) > 0 {
						auth.Status = e.SUCCESS
				} else {
						auth.Status = e.ERROR_AUTH_INACTIVE
				}
		}
	}
	return
}