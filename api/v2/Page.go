package api_v2

import (
	"github.com/gin-gonic/gin"

	"goserver/middlewares"
	models "goserver/models/v2"
)

// 查询列表参数
type GetPageListParams struct {
	Title string           `form:"title"`
	Path  string           `form:"path"`
	Tags  models.FlatArray `gorm:"TYPE:json" form:"tags"`
}

// 查询列表返回
type GetPageListResponse []struct {
	ID    string           `json:"id"`
	Title string           `json:"title"`
	Path  string           `json:"path"`
	Tags  models.FlatArray `gorm:"TYPE:json" json:"tags"`
}

func GetPageListApi(c *gin.Context) {
	middlewares.GetList(c, &GetPageListParams{}, &GetPageListResponse{}, &models.Page{})
}

// 查询页面参数
type GetPageParams struct {
	ID string `form:"id"`
}

// 查询页面返回
type GetPageResponse struct {
	models.Page
}

func GetPageApi(c *gin.Context) {
	middlewares.GetOne(c, &GetPageParams{}, &GetPageResponse{}, &models.Page{})
}

// 新增页面
func CreatePageApi(c *gin.Context) {
	middlewares.CreateOne(c, &models.Page{})
}

// 更新页面
func UpdatePageApi(c *gin.Context) {
	middlewares.UpdateOne(c, &models.Page{})
}

func DeletePageApi(c *gin.Context) {
	middlewares.DeleteOne(c, &models.Page{})
}

// 页面发布
type PublishPageParams struct {
	ID      string `json:"id"`
	Version string `json:"version"` // 发布的版本可回退，生成版本备份
}
