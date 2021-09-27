package models_v2

import (
	"goserver/libs/conf"
	"goserver/libs/gorm"
)

type User struct {
	Base
	Name string `json:"name"`
	RoleName  string    `json:"role_name"`
	AuthorityLevel int    `gorm:"default:1" json:"authority_level"`
	Email string `json:"email"`
	Password string `json:"password"`
	Valid int `gorm:"default:1" json:"valid"`
	LoginTime  string `gorm:"default:NOW()" json:"login_time"`
	LogoutTime string `gorm:"default:NOW()" json:"logout_time"`
}

var appUrl = conf.GetSectionKey("app", "APP_URL").String()

func GetUserById(id string) (user User) {
	gorm.GetDB().Where("id = ?", id).Find(&user)
	return;
}

func (user User) MakeConfirmationLink(confirmation string) string {
    return appUrl + "/api/confirm-signup?email=" + user.Email + "&code=" + confirmation
}

func (user User) MakePasswordResetLink(code string) string {
    return appUrl + "/reset-password?email=" + user.Email + "&code=" + code
}

func (user User) LogUserPersistence(persistence string) {
    p := Persistence{Base{ID: persistence}, user.ID}
    gorm.GetDB().Create(&p)
}