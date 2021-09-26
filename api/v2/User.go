package apiv2

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"goserver/libs/e"
	"goserver/libs/jwt"
	"goserver/libs/db"
	"goserver/libs/utils"
	"goserver/libs/mail"
	models "goserver/models/v2"
)

type Auth struct {
	User models.User
	Activation models.Activation
	Status int
}

func checkAuth(userName, password, captchaID, captchaCode string) (auth Auth) {
	auth.Status = e.ERROR_AUTH
	var user models.User
	db.DB().Where(models.User{Name: userName, Password: utils.EncryptPassword(password)}).First(&user)
	if len(user.ID) > 0 {
			var activation models.Activation
			db.DB().Where(models.Activation{UserId: user.ID}).Where("completed_at IS NOT NULL").First(&activation)
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

// 登录传参
type LoginParams struct {
	UserName    string `form:"user_name" json:"user_name"`
	Password    string `form:"password" json:"password"`
	CaptchaID   string `form:"captcha_id" json:"captcha_id" binding:"required"`     // 验证码ID
	CaptchaCode string `form:"captcha_code" json:"captcha_code" binding:"required"` // 验证码
}

// 登录返值
type LoginResponse struct {
	User      models.User `json:"user"`
	Token     string    `json:"token"`
	ExpiresAt int64     `json:"expiresAt"`
}

// 登录接口
func LoginApi(c *gin.Context) {
	var params LoginParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := make(map[string]interface{})
	code := e.ERROR_AUTH
	auth := checkAuth(params.UserName, params.Password, params.CaptchaID, params.CaptchaCode)
	if auth.Status == e.SUCCESS {
		persistence := utils.GenerateUuid()
		token, err := jwt.GenerateToken(auth.User.ID, auth.User.Name + " " + auth.User.RoleName, persistence)
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

	status := http.StatusBadRequest
	if code == e.SUCCESS {
		status = http.StatusOK
	}
	c.JSON(status, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// 退出登录接口
func LogoutApi(c *gin.Context) {
	// TODO
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out!"})
}


type ResetPasswordParams struct {
	// ID string `form:"id" json:"id" xml:"id" binding:"required"`
	UserName string `form:"user_name" json:"user_name" xml:"user_name" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" xml:"confirm_password" binding:"required"`
}

// 重置密码接口
func ResetPasswordApi(c *gin.Context) {
	var params ResetPasswordParams

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error: invalid data."})
		return
	}
	if params.ConfirmPassword != params.Password {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error: unmatched password."})
		return
	}

	var user models.User
	db.DB().Where(models.User{Name: params.UserName}).Where("deleted_at IS NULL").First(&user)
	if len(user.ID) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error: invalid user."})
		return
	}

	var activation models.Activation
	db.DB().Where(models.Activation{UserId: user.ID}).First(&activation)
	if len(activation.ID) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error: invalid activation."})
		return
	}
	if len(activation.CompletedAt) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error: You have reset your password."})
		return
	}

	db.DB().Exec("UPDATE user SET password=? WHERE id = ?", utils.EncryptPassword(params.Password), user.ID)
	db.DB().Exec("UPDATE activation SET completed_at=? WHERE id = ?", time.Now().Format("2006-01-02 15:04:05"), activation.ID)
	c.JSON(http.StatusOK, gin.H{"message": "Successfully reset password!" })
}

type SignupParams struct {
	UserName    string `form:"user_name" json:"user_name"`
	Email           string `form:"email" json:"email" xml:"email" binding:"required,email"`
	Phone           string `form:"phone" json:"phone" xml:"phone"`
	Password        string `form:"password" json:"password" xml:"password" binding:"required,min=6"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" xml:"confirm_password" binding:"required,min=6"`
}

func SignupApi(c *gin.Context) {
	var params SignupParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if params.Password != params.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password not matched"})
		return
	}

	var user models.User
	db.DB().Where(models.User{Email: params.Email}).Where("deleted_at IS NULL").First(&user)
	if len(user.ID) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already registered."})
		return
	}

	user = models.User{Base: models.Base{ID: utils.GenerateUuid()}, Name: params.UserName, Email: params.Email, Password: utils.EncryptPassword(params.Password)}
	db.Create(&user)
	activation := models.Activation{Base: models.Base{ID: utils.GenerateUuid()}, UserId: user.ID}
	db.Create(&activation)

	c.JSON(http.StatusOK, gin.H{
		"message": "You have signed up successfully. Please check you email for instructions to confirm your email address.",
	})

	mail.SendWelcomeEmail(params.Email, user.MakeConfirmationLink(activation.ID))
}