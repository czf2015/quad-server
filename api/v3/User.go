package api_v3

import (
	"goserver/libs/e"
	"goserver/libs/email"
	"goserver/libs/jwt"
	"goserver/libs/utils"
	models "goserver/models/v3"
	"net/http"
	"strconv"

	"goserver/middlewares"

	"github.com/gin-gonic/gin"
)

// func Register(c *gin.Context) {
// 	var user models.User
// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		// 处理请求参数错误
// 		return
// 	}

// 	// 在这里对密码进行加密处理
// 	// ...

// 	// 生成激活码
// 	activationCode := utils.GenerateUuid()
// 	user.ActivationCode = activationCode

// 	if err := db.Create(&user).Error; err != nil {
// 		// 处理数据库插入错误
// 		return
// 	}

// 	// 发送激活邮件等操作
// 	// ...

// 	c.JSON(http.StatusOK, gin.H{"message": "注册成功，请查收激活邮件"})
// }

// func Activate(c *gin.Context) {
// 	activationCode := c.Query("code")

// 	var user models.User
// 	db := db
// 	if err := db.Where("activation_code = ?", activationCode).First(&user).Error; err != nil {
// 		// 处理未找到用户或其他错误
// 		return
// 	}

// 	// 更新用户的激活状态
// 	user.Activated = true
// 	if err := db.Save(&user).Error; err != nil {
// 		// 处理数据库更新错误
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "账号激活成功"})
// }

// func Login(c *gin.Context) {
// 	var loginInfo struct {
// 		Email    string
// 		Password string
// 	}
// 	if err := c.ShouldBindJSON(&loginInfo); err != nil {
// 		// 处理请求参数错误
// 		return
// 	}

// 	var user models.User
// 	db := db
// 	if err := db.Where("email = ?", loginInfo.Email).First(&user).Error; err != nil {
// 		// 处理用户不存在错误
// 		return
// 	}

// 	// 验证密码
// 	if !CheckPassword(loginInfo.Password, user.PasswordHash) {
// 		// 处理密码错误
// 		return
// 	}

// 	// 检查账号激活状态
// 	if !user.Activated {
// 		// 处理账号未激活错误
// 		return
// 	}

// 	// 生成 JWT 等认证信息
// 	token, err := GenerateToken(user.ID)
// 	if err != nil {
// 		// 处理生成 token 错误
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"token": token})
// }

// 登录传参
type LoginParams struct {
	Name        string `form:"name" json:"name" binding:"required"`
	Password    string `form:"password" json:"password" binding:"required"`
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
	Name            string `form:"name" json:"name" xml:"name" binding:"required"`
	Password        string `form:"password" json:"password" xml:"password" binding:"required"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" xml:"confirm_password" binding:"required"`
}

type RegisterParams struct {
	Name string `form:"name" json:"name" binding:"required"`
	// RoleName string `form:"role_name" json:"role_name"`
	Email string `form:"email" json:"email" xml:"email" binding:"required"`
	// Phone           string `form:"phone" json:"phone" xml:"phone"`
	Password        string `form:"password" json:"password" xml:"password" binding:"required,min=6"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" xml:"confirm_password" binding:"required,min=6"`
}

// 用户注册接口
func Register(c *gin.Context) {
	var params RegisterParams
	code := e.SUCCESS
	if middlewares.BindJSON(c, &params) {
		if params.Password != params.ConfirmPassword {
			code = e.ERROR_UNMATCHED_PASSWORD
			c.JSON(http.StatusBadRequest, gin.H{"code": code, "message": e.GetMsg(code)})
			return
		}

		var user models.User
		if err := db.Where(models.User{Email: params.Email}).Where("deleted_at IS NULL").First(&user).Error; err != nil {
			// code = e.ERROR_ORM_GET
			// c.JSON(http.StatusInternalServerError, gin.H{"code": code, "message": e.GetMsg(code)})
			// return
		} else {
			if user.ID > 0 {
				code = e.ERROR_EMAIL_REGISTERED
				c.JSON(http.StatusInternalServerError, gin.H{"code": code, "message": e.GetMsg(code)})
				return
			}
		}

		user = models.User{Name: params.Name /* , RoleName: params.RoleName */, Email: params.Email, Password: utils.EncryptPassword(params.Password)}
		db.Create(&user)

		subject := "Welcome to lcdp"
		body := "<p>You have signed up successfully.</p>" +
			"<p>Please click the following link to confirm your email and activate your account.</p>" +
			"<p>" + user.MakeConfirmationLink(user.ActivationCode) + "</p>"
		go email.Send(params.Email, subject, body, "text/html")

		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": "You have signed up successfully. Please check you email for instructions to confirm your email address.",
		})
	}
}

// 登录接口
func Login(c *gin.Context) {
	var params LoginParams
	if middlewares.BindJSON(c, &params) {
		data := make(map[string]interface{})
		code := e.ERROR_AUTH
		status := http.StatusBadRequest
		auth := middlewares.CheckAuth(params.Name, params.Password, params.CaptchaID, params.CaptchaCode)
		if auth.Status == e.SUCCESS {
			persistence := utils.GenerateUuid()
			if token, err := jwt.GenerateToken(strconv.Itoa(int(auth.User.ID)), auth.User.Name /* +" "+auth.User.RoleName */, persistence); err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				auth.User.LogUserPersistence(persistence)
				data["token"] = token
				code = e.SUCCESS
				status = http.StatusOK
			}
		} else if auth.Status == e.ERROR_AUTH_INACTIVE {
			code = e.ERROR_AUTH_INACTIVE
		}

		c.JSON(status, gin.H{
			"code":    code,
			"message": e.GetMsg(code),
			"data":    data,
		})
	}
}

// 退出登录接口
func Logout(c *gin.Context) {
	code := e.SUCCESS
	token := c.GetHeader("Authorization")
	if token == "" {
		code = e.INVALID_PARAMS
	} else {
		if claims, err := jwt.ParseToken(token); err != nil {
			code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
		} else {
			token := jwt.ExpireToken(*claims)
			c.Header("Authorization", token)
		}
	}

	if code != e.SUCCESS {
		c.JSON(http.StatusInternalServerError, gin.H{"code": code, "message": e.GetMsg(code)})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": code, "message": "Successfully logged out!"})
	}
}

// 重置密码接口
func ResetPassword(c *gin.Context) {
	var params ResetPasswordParams
	code := e.ERROR
	if middlewares.BindJSON(c, &params) {
		if params.ConfirmPassword != params.Password {
			c.JSON(http.StatusBadRequest, gin.H{"code": code, "message": e.GetMsg(code)})
			return
		}

		var user models.User
		db.Where(models.User{Name: params.Name}).Where("deleted_at IS NULL").First(&user)
		if user.ID == 0 {
			code = e.ERROR_INVALID_USER
			c.JSON(http.StatusBadRequest, gin.H{"code": code, "message": e.GetMsg(code)})
			return
		}

		db.Exec("UPDATE user SET password=?, activated=? WHERE id = ?", utils.EncryptPassword(params.Password), true, user.ID)
		c.JSON(http.StatusOK, gin.H{"code": code, "message": "Successfully reset password!"})
	}
}

func CreateUser(c *gin.Context) {
	middlewares.Create(c, &models.User{})
}

func GetUser(c *gin.Context) {
	middlewares.Get(c, &models.User{})
}

func UpdateUser(c *gin.Context) {
	middlewares.Update(c, &models.User{})
}

func DeleteUser(c *gin.Context) {
	middlewares.Delete(c, &models.User{})
}
