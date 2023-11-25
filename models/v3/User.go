package models_v3

import (
	"goserver/libs/gorm"
	"goserver/libs/utils"
	"time"
)

type User struct {
	Base
	Name           string     `gorm:"type:varchar(255);not null" json:"name"`
	Email          string     `gorm:"type:varchar(255);not null" json:"email"`
	Company        string     `gorm:"type:varchar(255)" json:"company"`
	Password       string     `gorm:"type:varchar(255);not null" json:"password"`
	Activated      bool       `json:"-"`
	ActivationCode string     `gorm:"type:varchar(255);not null" json:"-"`
	Persistence    string     `gorm:"type:varchar(255);not null" json:"-"`
	Roles          FlatArray  `json:"roles"`
	LoginTime      *time.Time `gorm:"type:timestamp" json:"login_time"`
	LogoutTime     *time.Time `gorm:"type:timestamp" json:"logout_time"`
}

func (user *User) BeforeCreate(db *gorm.DB) (err error) {
	user.ActivationCode = utils.GenerateUuid()

	return
}

func GetUserById(id uint) (user User) {
	db.Where("id = ?", id).Find(&user)
	return
}

func (user User) MakeConfirmationLink(activation_id string) string {
	return appUrl + "/api/confirm-signup?email=" + user.Email + "&code=" + activation_id
}

func (user User) MakePasswordResetLink(code string) string {
	return appUrl + "/reset-password?email=" + user.Email + "&code=" + code
}

func (user User) LogUserPersistence(persistence string) {
	// p := Persistence{ID: persistence_id, UserId: user.ID}
	user.Persistence = persistence
	db.Save(&user)
}
