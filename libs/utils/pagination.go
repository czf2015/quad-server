package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"goserver/libs/conf"
)

var pageSize = conf.GetSectionKey("app", "PAGE_SIZE").MustInt()

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * pageSize
	}

	return result
}