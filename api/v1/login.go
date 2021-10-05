package api_v1

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"goserver/libs/e"
	"goserver/libs/gorm"
	"goserver/libs/jwt"
	"goserver/libs/mail"
	"goserver/libs/utils"
	"goserver/models"
)

const (
	userkey = "user"
)

type Account struct {
	Email    string `form:"email" json:"email" xml:"email" binding:"required,email"`
	Password string `form:"password" json:"password" xml:"password" binding:"required,min=6"`
}

type SendPasswordResetJson struct {
	Email string `form:"email" json:"email" xml:"email" binding:"required,email"`
}

type PasswordResetJson struct {
	Email           string `form:"email" json:"email" xml:"email" binding:"required,email"`
	Code            string `form:"code" json:"code" xml:"code" binding:"required"`
	Password        string `form:"password" json:"password" xml:"password" binding:"required"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" xml:"confirm_password" binding:"required"`
}

func LoginApi(c *gin.Context) {
	var json Account
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := make(map[string]interface{})
	code := e.ERROR_AUTH
	auth := models.CheckAuth(json.Email, json.Password)
	if auth.Status == e.SUCCESS {
		persistence := utils.GenerateUuid()
		token, err := jwt.GenerateToken(auth.User.ID, auth.User.FirstName+" "+auth.User.LastName, persistence)
		if err != nil {
			code = e.ERROR_AUTH_TOKEN
		} else {
			auth.User.LogUserPersistence(persistence)
			data["token"] = token
			code = e.SUCCESS
		}
	} else if auth.Status == e.ERROR_AUTH_INACTIVE {
		code = e.ERROR_AUTH_INACTIVE
	}

	if code != e.SUCCESS {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
	}
}

func LogoutApi(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

func LoginStatusApi(c *gin.Context) {
	claims, _ := jwt.ParseToken(c.Query("token"))
	refreshToken := jwt.RefreshToken(*claims)
	c.JSON(http.StatusOK, gin.H{"status": true, "name": claims.Name, "persistence": claims.Persistence, "refreshToken": refreshToken})
}

func SendResetPasswordApi(c *gin.Context) {
	var resetJson SendPasswordResetJson
	if err := c.ShouldBindJSON(&resetJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var user models.User
	gorm.GetDB().Where(models.User{Email: resetJson.Email}).Where("deleted_at IS NULL").First(&user)
	if len(user.ID) > 0 {
		activation := models.Activation{Base: models.Base{ID: utils.GenerateUuid()}, UserId: user.ID}
		gorm.Create(&activation)
		mail.SendResetPasswordEmail(user.Email, user.MakePasswordResetLink(activation.ID))
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Please check your email."})
}

func ResetPasswordApi(c *gin.Context) {
	var resetJson PasswordResetJson
	if err := c.ShouldBindJSON(&resetJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error: invalid data."})
		return
	}
	if resetJson.ConfirmPassword != resetJson.Password {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error: unmatched password."})
		return
	}
	var user models.User
	gorm.GetDB().Where(models.User{Email: resetJson.Email}).Where("deleted_at IS NULL").First(&user)
	if len(user.ID) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error: invalid email."})
		return
	}
	var activation models.Activation
	gorm.GetDB().Where(models.Activation{Base: models.Base{ID: resetJson.Code}, UserId: user.ID}).First(&activation)
	if len(activation.ID) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error: invalid code."})
		return
	}
	if len(activation.CompletedAt) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "You have reset your password."})
		return
	}
	gorm.GetDB().Exec("UPDATE user SET password=? WHERE id = ?", utils.EncryptPassword(resetJson.Password), user.ID)
	gorm.GetDB().Exec("UPDATE activation SET completed_at=? WHERE id = ?", time.Now().Format("2006-01-02 15:04:05"), activation.ID)
	c.JSON(http.StatusOK, gin.H{"message": "Success."})
}

func ConfirmSignUpApi(c *gin.Context) {
	email := c.Query("email")
	code := c.Query("code")
	if len(email) == 0 || len(code) == 0 {
		c.HTML(http.StatusOK, "redirect.tmpl", gin.H{"message": "Error.", "url": "/", "seconds": 3})
		return
	}
	var user models.User
	gorm.GetDB().Where(models.User{Email: email}).Where("deleted_at IS NULL").First(&user)
	if len(user.ID) == 0 {
		c.HTML(http.StatusOK, "redirect.tmpl", gin.H{"message": "Error.", "url": "/", "seconds": 3})
		return
	}
	var activation models.Activation
	gorm.GetDB().Where(models.Activation{Base: models.Base{ID: code}}).First(&activation)
	if len(activation.ID) == 0 {
		c.HTML(http.StatusOK, "redirect.tmpl", gin.H{"message": "Error.", "url": "/", "seconds": 3})
		return
	}
	if len(activation.CompletedAt) > 0 {
		c.HTML(http.StatusOK, "redirect.tmpl", gin.H{"message": "You are already confirmed.", "url": "/", "seconds": 3})
		return
	}
	gorm.GetDB().Model(&activation).Update("completed_at", time.Now().Format("2006-01-02 15:04:05"))
	c.HTML(http.StatusOK, "redirect.tmpl", gin.H{"message": "You have been successfully confirmed.", "url": "/", "seconds": 3})
	return
}
