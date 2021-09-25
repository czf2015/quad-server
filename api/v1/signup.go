package apiv1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"goserver/models"
	"goserver/libs/db"
	"goserver/libs/utils"
	"goserver/libs/mail"
)

type Signature struct {
	FirstName       string `form:"first_name" json:"first_name" xml:"first_name"  binding:"required"`
	LastName        string `form:"last_name" json:"last_name" xml:"last_name"  binding:"required"`
	Email           string `form:"email" json:"email" xml:"email" binding:"required,email"`
	Password        string `form:"password" json:"password" xml:"password" binding:"required,min=6"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" xml:"confirm_password" binding:"required,min=6"`
	Phone           string `form:"phone" json:"phone" xml:"phone"`
	Website         string `form:"website" json:"website" xml:"website"`
}

func SignupApi(c *gin.Context) {
	var json Signature
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if json.Password != json.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password not matched"})
		return
	}

	var exist models.User
	db.DB().Where(models.User{Email: json.Email}).Where("deleted_at IS NULL").First(&exist)
	if len(exist.ID) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists."})
		return
	}

	user := models.User{Base: models.Base{ID: utils.GenerateUuid()}, FirstName: json.FirstName, LastName: json.LastName, Email: json.Email, Password: utils.EncryptPassword(json.Password), Phone: json.Phone, Website: json.Website}
	db.Create(&user)
	activation := models.Activation{Base: models.Base{ID: utils.GenerateUuid()}, UserId: user.ID}
	db.Create(&activation)

	c.JSON(200, gin.H{
		"message": "You have signed up successfully. Please check you email for instructions to confirm your email address.",
	})

	mail.SendWelcomeEmail(json.Email, user.MakeConfirmationLink(activation.ID))
}