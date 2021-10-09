package api_v2

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"goserver/libs/e"
	"goserver/libs/gorm"
	"goserver/libs/jwt"
	"goserver/libs/mail"
	"goserver/libs/utils"
	"goserver/middlewares"
	models "goserver/models/v2"
)

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
	Token     string      `json:"token"`
	ExpiresAt int64       `json:"expiresAt"`
}

type ResetPasswordParams struct {
	// ID string `form:"id" json:"id" xml:"id" binding:"required"`
	UserName        string `form:"user_name" json:"user_name" xml:"user_name" binding:"required"`
	Password        string `form:"password" json:"password" xml:"password" binding:"required"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" xml:"confirm_password" binding:"required"`
}

type SignupParams struct {
	UserName        string `form:"user_name" json:"user_name" binding:"required"`
	RoleName        string `form:"role_name" json:"role_name"`
	Email           string `form:"email" json:"email" xml:"email" binding:"required"`
	Phone           string `form:"phone" json:"phone" xml:"phone"`
	Password        string `form:"password" json:"password" xml:"password" binding:"required,min=6"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" xml:"confirm_password" binding:"required,min=6"`
}

// 登录接口
func LoginApi(c *gin.Context) {
	var params LoginParams
	if middlewares.BindJSON(c, &params) {
		data := make(map[string]interface{})
		code := e.ERROR_AUTH
		auth := middlewares.CheckAuth(params.UserName, params.Password, params.CaptchaID, params.CaptchaCode)
		if auth.Status == e.SUCCESS {
			persistence := utils.GenerateUuid()
			token, err := jwt.GenerateToken(auth.User.ID, auth.User.Name+" "+auth.User.RoleName, persistence)
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
}

// 退出登录接口
func LogoutApi(c *gin.Context) {
	// TODO
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out!"})
}

// 重置密码接口
func ResetPasswordApi(c *gin.Context) {
	var params ResetPasswordParams
	if middlewares.BindJSON(c, &params) {
		if params.ConfirmPassword != params.Password {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Error: unmatched password."})
			return
		}

		var user models.User
		gorm.GetDB().Where(models.User{Name: params.UserName}).Where("deleted_at IS NULL").First(&user)
		if len(user.ID) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Error: invalid user."})
			return
		}

		var activation models.Activation
		gorm.GetDB().Where(models.Activation{UserId: user.ID}).First(&activation)
		if len(activation.ID) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Error: invalid activation."})
			return
		}
		if len(activation.CompletedAt) > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Error: You have reset your password."})
			return
		}

		gorm.GetDB().Exec("UPDATE user SET password=? WHERE id = ?", utils.EncryptPassword(params.Password), user.ID)
		gorm.GetDB().Exec("UPDATE activation SET completed_at=? WHERE id = ?", time.Now().Format("2006-01-02 15:04:05"), activation.ID)
		c.JSON(http.StatusOK, gin.H{"message": "Successfully reset password!"})
	}
}

// 用户注册接口
func SignupApi(c *gin.Context) {
	var params SignupParams
	if middlewares.BindJSON(c, &params) {
		if params.Password != params.ConfirmPassword {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Password not matched"})
			return
		}

		var user models.User
		gorm.GetDB().Where(models.User{Email: params.Email}).Where("deleted_at IS NULL").First(&user)
		if len(user.ID) > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email already registered."})
			return
		}

		user = models.User{Base: models.Base{ID: utils.GenerateUuid()}, Name: params.UserName, RoleName: params.RoleName, Email: params.Email, Password: utils.EncryptPassword(params.Password)}
		gorm.Create(&user)
		activation := models.Activation{Base: models.Base{ID: utils.GenerateUuid()}, UserId: user.ID}
		gorm.Create(&activation)

		c.JSON(http.StatusOK, gin.H{
			"message": "You have signed up successfully. Please check you email for instructions to confirm your email address.",
		})

		mail.SendWelcomeEmail(params.Email, user.MakeConfirmationLink(activation.ID))
	}
}
