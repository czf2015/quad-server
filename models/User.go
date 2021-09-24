package models

import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
    
	"github.com/google/uuid"
	"github.com/thoas/go-funk"

    "goserver/libs/conf"
    "goserver/libs/e"
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

type Activation struct {
    Base
    UserId string `json:"user_id"`
    CompletedAt string `gorm:"default:NULL" json:"completed_at"`
}

type Persistence struct {
    Base
    UserId string `json:"user_id"`
}

type Auth struct {
    User User
    Activation Activation
    Status int
}


func EncryptPassword(password string) string {
    // Create a new HMAC by defining the hash type and the key (as byte array)
    h := hmac.New(sha256.New, []byte(conf.GetSectionKey("app", "APP_KEY").String()))
    // Write Data to it
    h.Write([]byte(password))
    // Get result and encode as hexadecimal string
    return hex.EncodeToString(h.Sum(nil))
}

func GenerateUuid() string {
    return uuid.New().String()
}

func CheckAuth(email, password string) (auth Auth) {
    auth.Status = e.ERROR_AUTH
    var user User
    db.DB().Where(User{Email: email, Password: EncryptPassword(password)}).First(&user)
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

func MakeConfirmationLink(user User, confirmation string) string {
    return conf.GetSectionKey("app", "APP_URL").String() + "/api/confirm-signup?email=" + user.Email + "&code=" + confirmation
}

func MakePasswordResetLink(user User, code string) string {
    return conf.GetSectionKey("app", "APP_URL").String() + "/reset-password?email=" + user.Email + "&code=" + code
}

func LogUserPersistence(user User, persistence string) {
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

func GetUserById(id string) (user User) {
	db.DB().Where("id = ?", id).Find(&user)
	return;
}