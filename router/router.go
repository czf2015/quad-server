package router

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"goserver/libs/conf"
	"goserver/middlewares"

	v1 "goserver/api/v1"
	v2 "goserver/api/v2"
)

var Router *gin.Engine

func init()  {
	gin.SetMode(conf.GetSectionKey("base", "RUN_MODE").String())

	Router = gin.New()

	
	store := sessions.NewCookieStore([]byte("secret"))
	Router.Use(sessions.Sessions("mysession", store))
	
	Router.Use(gin.Logger())
	
	Router.Use(gin.Recovery())
	
	Router.LoadHTMLGlob("templates/*")

	apiv1 := Router.Group("/api/v1")
	{
		apiv1.POST("/login", v1.Login)
		apiv1.GET("/logout", v1.Logout)
		apiv1.POST("/signup", v1.Signup)
		apiv1.POST("/send-reset-password", v1.SendResetPassword)
		apiv1.POST("/reset-password", v1.ResetPassword)
		apiv1.POST("/contactus", v1.Contactus)
		apiv1.GET("/confirm-signup", v1.ConfirmSignUp)
	}
	apiv1.Use(middlewares.JWT())
	{
		apiv1.GET("/login-status", v1.LoginStatus)
		apiv1.GET("/report/last-30-days", v1.GetLast30DaysReport)
		apiv1.GET("/agreement/unsigned", v1.UnsignedAgreements)
		apiv1.POST("/agreement/sign", v1.SignAgreement)
	}

	apiv2 := Router.Group("/api/v2")
	{
		apiv2.GET("/captcha", v2.GetCaptchaApi)
	}
}