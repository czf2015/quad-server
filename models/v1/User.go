package models_v1

import (    
	"github.com/thoas/go-funk"

    "goserver/libs/conf"
    "goserver/libs/gorm"
)

type User struct {
    Base

    FirstName string `json:"first_name"`
    LastName string `json:"last_name"`
    Email string `json:"email"`
    Password string `json:"password"`
    Phone string `json:"phone"`
	Website string `json:"website"`
	
	Agreements []Agreement `gorm:"many2many:user_agreement;"`
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

func (user User) GetApprovedDomains() (approvedDomains []ApprovedDomain) {
    gorm.GetDB().Where(ApprovedDomain{UserId: user.ID, Approved: true}).Find(&approvedDomains)
    return
}

func (user User) GetLatestUnsignedAgreements() (unsigned []Agreement) {
    var agreements []Agreement
    gorm.GetDB().Model(&user).Related(&agreements, "Agreements")
    ids := funk.Map(agreements, func (agreement Agreement) string {
        return agreement.ID
    }).([]string)
    latest := GetLatestAgreements()
    latestIds := funk.Map(latest, func (agreement Agreement) string {
        return agreement.ID
    }).([]string)
    _, unsignedIds := funk.DifferenceString(ids, latestIds)
    unsigned = funk.Filter(latest, func (agreement Agreement) bool {
        return funk.Contains(unsignedIds, agreement.ID)
    }).([]Agreement)
	return unsigned
}