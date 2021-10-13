package api_v2

import (
	"github.com/gin-gonic/gin"

	"goserver/libs/crawl"
	"goserver/middlewares"
	models "goserver/models/v2"
)


func Crawl(c *gin.Context) {
	var params models.Link
	if middlewares.BindQuery(c, &params) {
		crawl.Crawl(params.Url)
	}
}