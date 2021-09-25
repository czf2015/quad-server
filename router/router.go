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
	runMode := conf.GetSectionKey("base", "RUN_MODE").String()
	gin.SetMode(runMode)

	Router = gin.New()

	store := sessions.NewCookieStore([]byte("secret"))
	Router.Use(sessions.Sessions("mysession", store))
	
	Router.Use(gin.Logger())
	
	Router.Use(gin.Recovery())
	
	Router.LoadHTMLGlob("templates/*")

	apiv1 := Router.Group("/api/v1")
	{
		apiv1.POST("/login", v1.LoginApi)
		apiv1.GET("/logout", v1.LogoutApi)
		apiv1.POST("/signup", v1.SignupApi)
		apiv1.POST("/send-reset-password", v1.SendResetPasswordApi)
		apiv1.POST("/reset-password", v1.ResetPasswordApi)
		apiv1.POST("/contactus", v1.ContactusApi)
		apiv1.GET("/confirm-signup", v1.ConfirmSignUpApi)
	}
	apiv1.Use(middlewares.JWT())
	{
		apiv1.GET("/login-status", v1.LoginStatusApi)
		apiv1.GET("/report/last-30-days", v1.GetLast30DaysReportApi)
		apiv1.GET("/agreement/unsigned", v1.UnsignedAgreementsApi)
		apiv1.POST("/agreement/sign", v1.SignAgreementApi)
	}

	apiv2 := Router.Group("/api/v2")
	{
		apiv2.GET("/captcha", v2.GetCaptchaApi)
	}
}