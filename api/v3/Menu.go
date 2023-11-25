package api_v3

import (
	models "goserver/models/v3"

	"goserver/middlewares"

	"github.com/gin-gonic/gin"
)

func CreateMenu(c *gin.Context) {
	middlewares.Create(c, &models.Menu{})
}

func GetMenu(c *gin.Context) {
	middlewares.Get(c, &models.Menu{})
}

func UpdateMenu(c *gin.Context) {
	middlewares.Update(c, &models.Menu{})
}

func DeleteMenu(c *gin.Context) {
	middlewares.Delete(c, &models.Menu{})
}

// 查询列表参数
type GetMenuListParams struct {
	Title string `form:"title"`
	Link  string `form:"link"`
}

// 查询列表返回
type GetMenuListResponse []struct {
	ID    uint   `json:"id"`
	PID   uint   `json:"pid"`
	Title string `json:"title"`
	Link  string `json:"link"`
	Order int    `json:"order"`
}

func GetMenuList(c *gin.Context) {
	middlewares.GetList(c, &GetMenuListParams{}, &GetMenuListResponse{}, &models.Menu{})
}
