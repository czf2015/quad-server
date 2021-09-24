package models

import (    
	"goserver/libs/e"
	"goserver/libs/db"
	"goserver/libs/utils"
)

type Auth struct {
	User User
	Activation Activation
	Status int
}

func CheckAuth(email, password string) (auth Auth) {
	auth.Status = e.ERROR_AUTH
	var user User
	db.DB().Where(User{Email: email, Password: utils.EncryptPassword(password)}).First(&user)
	if len(user.ID) > 0 {
			var activation Activation
			db.DB().Where(Activation{UserId: user.ID}).Where("completed_at IS NOT NULL").First(&activation)
			auth.User = user
			auth.Activation = activation
			if len(activation.CompletedAt) > 0 {
					auth.Status = e.SUCCESS
			} else {
					auth.Status = e.ERROR_AUTH_INACTIVE
			}
	}
	return
}