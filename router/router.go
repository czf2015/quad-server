package router

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"goserver/libs/conf"

	// v1 "goserver/api/v1"

	v3 "goserver/api/v3"
)

var Router *gin.Engine

func init() {
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

	// apiv2 := Router.Group("/api/v2")
	{
		// user
		// apiv2.GET("/captcha", v2.GetCaptchaApi)
		// apiv2.POST("/login", v2.LoginApi)
		// apiv2.GET("/logout", v2.LogoutApi)
		// apiv2.POST("/signup", v2.SignupApi)
		// apiv2.POST("/reset-password", v2.ResetPasswordApi)
		// // address
		// apiv2.POST("/address", v2.CreateAddressApi)
		// apiv2.POST("/address-list", v2.CreateAddressListApi)
		// apiv2.GET("/address", v2.GetAddressApi)
		// apiv2.GET("/address-all", v2.GetAllAddressApi)
		// apiv2.PUT("/address", v2.UpdateAddressApi)
		// apiv2.PUT("/address-list", v2.UpdateAddressListApi)
		// apiv2.DELETE("/address", v2.DeleteAddressApi)
		// apiv2.DELETE("/address-list", v2.DeleteAddressListApi)
		// // area
		// apiv2.POST("/area", v2.CreateAreaApi)
		// apiv2.POST("/area-list", v2.CreateAreaListApi)
		// apiv2.GET("/area", v2.GetAreaApi)
		// apiv2.GET("/area-all", v2.GetAllAreaApi)
		// apiv2.PUT("/area", v2.UpdateAreaApi)
		// apiv2.PUT("/area-list", v2.UpdateAreaListApi)
		// apiv2.DELETE("/area", v2.DeleteAreaApi)
		// apiv2.DELETE("/area-list", v2.DeleteAreaListApi)
		// // addressPlan
		// apiv2.POST("/address-plan", v2.CreateAddressPlanApi)
		// apiv2.POST("/address-plan-list", v2.CreateAddressPlanListApi)
		// apiv2.GET("/address-plan", v2.GetAddressPlanApi)
		// apiv2.GET("/address-plan-all", v2.GetAllAddressPlanApi)
		// apiv2.GET("/address-plan-list", v2.GetAddressPlanListApi)
		// apiv2.PUT("/address-plan", v2.UpdateAddressPlanApi)
		// apiv2.PUT("/address-plan-list", v2.UpdateAddressPlanListApi)
		// apiv2.DELETE("/address-plan", v2.DeleteAddressPlanApi)
		// apiv2.DELETE("/address-plan-list", v2.DeleteAddressPlanListApi)
		// // networkAllocation
		// apiv2.POST("/network-allocation", v2.CreateNetworkAllocationApi)
		// apiv2.POST("/network-allocation-list", v2.CreateNetworkAllocationListApi)
		// apiv2.GET("/network-allocation", v2.GetNetworkAllocationApi)
		// apiv2.GET("/network-allocation-all", v2.GetAllNetworkAllocationApi)
		// apiv2.GET("/network-allocation-list", v2.GetNetworkAllocationListApi)
		// apiv2.PUT("/network-allocation", v2.UpdateNetworkAllocationApi)
		// apiv2.PUT("/network-allocation-list", v2.UpdateNetworkAllocationListApi)
		// apiv2.DELETE("/network-allocation", v2.DeleteNetworkAllocationApi)
		// apiv2.DELETE("/network-allocation-list", v2.DeleteNetworkAllocationListApi)
		// // networkManage
		// apiv2.POST("/network-manage", v2.CreateNetworkManageApi)
		// apiv2.POST("/network-manage-list", v2.CreateNetworkManageListApi)
		// apiv2.GET("/network-manage", v2.GetNetworkManageApi)
		// apiv2.GET("/network-manage-all", v2.GetAllNetworkManageApi)
		// apiv2.GET("/network-manage-list", v2.GetNetworkManageListApi)
		// apiv2.PUT("/network-manage", v2.UpdateNetworkManageApi)
		// apiv2.PUT("/network-manage-list", v2.UpdateNetworkManageListApi)
		// apiv2.DELETE("/network-manage", v2.DeleteNetworkManageApi)
		// apiv2.DELETE("/network-manage-list", v2.DeleteNetworkManageListApi)
		// // page
		// apiv2.GET("/pages/list", v2.GetPageListApi)
		// apiv2.GET("/pages", v2.GetPageApi)
		// apiv2.POST("/pages", v2.CreatePageApi)
		// apiv2.PUT("/pages", v2.UpdatePageApi)
		// apiv2.DELETE("/pages", v2.DeletePageApi)
		// // publish
		// apiv2.GET("/pages/publish/list", v2.GetPublishListApi)
		// apiv2.GET("/pages/publish", v2.GetPublishApi)
		// apiv2.POST("/pages/publish", v2.CreatePublishApi)
		// apiv2.PUT("/pages/publish", v2.UpdatePublishApi)
		// apiv2.DELETE("/pages/publish", v2.DeletePageApi)

		// apiv2.GET("/third", third.GetThirdService)
		// apiv2.GET("/crawl", v2.Crawl)
	}

	apiv3 := Router.Group("/api/v3")
	{
		apiv3.POST("/register", v3.Register)
		apiv3.GET("/captcha", v3.GetCaptcha)
		apiv3.POST("/login", v3.Login)
	}
	// apiv3.Use(middlewares.JWT())
	{
		apiv3.POST("/logout", v3.Logout)
		apiv3.POST("/reset/password", v3.ResetPassword)

		// 路由定义
		apiv3.POST("/users", v3.CreateUser)
		apiv3.GET("/users/:id", v3.GetUser)
		apiv3.PUT("/users/:id", v3.UpdateUser)
		apiv3.DELETE("/users/:id", v3.DeleteUser)

		apiv3.POST("/roles", v3.CreateRole)
		apiv3.GET("/roles/:id", v3.GetRole)
		apiv3.PUT("/roles/:id", v3.UpdateRole)
		apiv3.DELETE("/roles/:id", v3.DeleteRole)

		apiv3.POST("/permissions", v3.CreatePermission)
		apiv3.GET("/permissions/:id", v3.GetPermission)
		apiv3.PUT("/permissions/:id", v3.UpdatePermission)
		apiv3.DELETE("/permissions/:id", v3.DeletePermission)

		// page
		apiv3.GET("/pages", v3.GetPageListApi)
		apiv3.GET("/pages/details", v3.GetPageApi)
		apiv3.POST("/pages", v3.CreatePageApi)
		apiv3.PUT("/pages/:id", v3.UpdatePageApi)
		apiv3.DELETE("/pages/:id", v3.DeletePageApi)
		apiv3.DELETE("/pages", v3.DeletePageListApi)
		// template
		apiv3.GET("/templates", v3.GetTemplateListApi)
		apiv3.GET("/templates/details", v3.GetTemplateApi)
		apiv3.POST("/templates", v3.CreateTemplateApi)
		apiv3.PUT("/templates/:id", v3.UpdateTemplateApi)
		apiv3.DELETE("/templates/:id", v3.DeleteTemplateApi)
		apiv3.DELETE("/templates", v3.DeleteTemplateListApi)
		// publish
		apiv3.GET("/publish", v3.GetPublishListApi)
		apiv3.GET("/publish/details", v3.GetPublishApi)
		apiv3.POST("/publish", v3.CreatePublishApi)
		apiv3.PUT("/publish/:id", v3.UpdatePublishApi)
		apiv3.PATCH("/publish/:id", v3.PatchPublishApi)
		apiv3.DELETE("/publish/:id", v3.DeletePublishApi)
		apiv3.DELETE("/publish", v3.DeletePublishListApi)

		apiv3.POST("/menus", v3.CreateMenu)
		apiv3.GET("/menus", v3.GetMenuList)
		apiv3.GET("/menus/:id", v3.GetMenu)
		apiv3.PUT("/menus/:id", v3.UpdateMenu)
		apiv3.DELETE("/menus/:id", v3.DeleteMenu)

		// event_rule
		apiv3.GET("/event/rules", v3.GetEventRuleListApi)
		apiv3.GET("/event/rules/details", v3.GetEventRuleApi)
		apiv3.POST("/event/rules", v3.CreateEventRuleApi)
		apiv3.PUT("/event/rules/:id", v3.UpdateEventRuleApi)
		apiv3.DELETE("/event/rules/:id", v3.DeleteEventRuleApi)
		apiv3.DELETE("/event/rules", v3.DeleteEventRuleListApi)
		// upload
		apiv3.POST("/upload", v3.UploadFile)
	}
}
