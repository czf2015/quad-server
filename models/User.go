package models

import (    
	"github.com/thoas/go-funk"

    "goserver/libs/conf"
    "goserver/libs/db"
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

func GetUserById(id string) (user User) {
	db.DB().Where("id = ?", id).Find(&user)
	return;
}

func (user User) MakeConfirmationLink(confirmation string) string {
    return conf.GetSectionKey("app", "APP_URL").String() + "/api/confirm-signup?email=" + user.Email + "&code=" + confirmation
}

func (user User) MakePasswordResetLink(code string) string {
    return conf.GetSectionKey("app", "APP_URL").String() + "/reset-password?email=" + user.Email + "&code=" + code
}

func (user User) LogUserPersistence(persistence string) {
    p := Persistence{Base{ID: persistence}, user.ID}
    db.DB().Create(&p)
}

func (user User) GetApprovedDomains() (approvedDomains []ApprovedDomain) {
    db.DB().Where(ApprovedDomain{UserId: user.ID, Approved: true}).Find(&approvedDomains)
    return
}

func (user User) GetLatestUnsignedAgreements() (unsigned []Agreement) {
    var agreements []Agreement
    db.DB().Model(&user).Related(&agreements, "Agreements")
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