package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"goserver/libs/conf"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	appCfg, _ := conf.GetSection("app")
	pageSize := appCfg.Key("PAGE_SIZE").MustInt()
	if page > 0 {
		result = (page - 1) * pageSize
	}

	return result
}