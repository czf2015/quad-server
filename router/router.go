package router

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"goserver/libs/conf"
	// "goserver/middlewares"

	// v1 "goserver/api/v1"
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

	Router.StaticFS("/static", http.Dir("./static"))

	// apiv1 := Router.Group("/api/v1")
	// {
	// 	apiv1.POST("/login", v1.LoginApi)
	// 	apiv1.GET("/logout", v1.LogoutApi)
	// 	apiv1.POST("/signup", v1.SignupApi)
	// 	apiv1.POST("/send-reset-password", v1.SendResetPasswordApi)
	// 	apiv1.POST("/reset-password", v1.ResetPasswordApi)
	// 	apiv1.POST("/contactus", v1.ContactusApi)
	// 	apiv1.GET("/confirm-signup", v1.ConfirmSignUpApi)
	// }
	// apiv1.Use(middlewares.JWT())
	// {
	// 	apiv1.GET("/login-status", v1.LoginStatusApi)
	// 	apiv1.GET("/report/last-30-days", v1.GetLast30DaysReportApi)
	// 	apiv1.GET("/agreement/unsigned", v1.UnsignedAgreementsApi)
	// 	apiv1.POST("/agreement/sign", v1.SignAgreementApi)
	// }

	apiv2 := Router.Group("/api/v2")
	{
		// user
		apiv2.GET("/captcha", v2.GetCaptchaApi)
		apiv2.POST("/login", v2.LoginApi)
		apiv2.GET("/logout", v2.LogoutApi)
		apiv2.POST("/signup", v2.SignupApi)
		apiv2.POST("/reset-password", v2.ResetPasswordApi)
		// address
		apiv2.POST("/address", v2.CreateAddressApi)
		apiv2.POST("/address-list", v2.CreateAddressListApi)
		apiv2.GET("/address", v2.GetAddressApi)
		apiv2.GET("/address-all", v2.GetAllAddressApi)
		apiv2.PUT("/address", v2.UpdateAddressApi)
		apiv2.PUT("/address-list", v2.UpdateAddressListApi)
		apiv2.DELETE("/address", v2.DeleteAddressApi)
		apiv2.DELETE("/address-list", v2.DeleteAddressListApi)
		// area
		apiv2.POST("/area", v2.CreateAreaApi)
		apiv2.POST("/area-list", v2.CreateAreaListApi)
		apiv2.GET("/area", v2.GetAreaApi)
		apiv2.GET("/area-all", v2.GetAllAreaApi)
		apiv2.PUT("/area", v2.UpdateAreaApi)
		apiv2.PUT("/area-list", v2.UpdateAreaListApi)
		apiv2.DELETE("/area", v2.DeleteAreaApi)
		apiv2.DELETE("/area-list", v2.DeleteAreaListApi)
	}
}