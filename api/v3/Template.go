package api_v3

import (
	"time"

	"github.com/gin-gonic/gin"

	"goserver/middlewares"
	models "goserver/models/v3"
)

// 查询列表参数
type GetTemplateListParams struct {
	Title string `form:"title"`
	Path  string `form:"path"`
	// Tags  []string `form:"tags"`
}

// 查询列表返回
type GetTemplateListResponse []struct {
	ID        uint             `json:"id"`
	Title     string           `json:"title"`
	Path      string           `json:"path"`
	Tags      models.FlatArray `gorm:"TYPE:json" json:"tags"`
	ImgUrl    string           `gorm:"type:varchar(255);not null;column:img_url" json:"imgUrl"`
	CreatedAt *time.Time       `gorm:"type:timestamp;default:NOW()" json:"create_time"`
	UpdatedAt *time.Time       `gorm:"type:timestamp;default:NOW()" json:"update_time"`
}

func GetTemplateListApi(c *gin.Context) {
	middlewares.GetList(c, &GetTemplateListParams{}, &GetTemplateListResponse{}, &models.Template{})
}

// 查询页面参数
type GetTemplateParams struct {
	ID uint `form:"id"`
}

// 查询页面返回
type GetTemplateResponse struct {
	models.Template
}

func GetTemplateApi(c *gin.Context) {
	middlewares.GetOne(c, &GetTemplateParams{}, &GetTemplateResponse{}, &models.Template{})
}

// 新增页面
func CreateTemplateApi(c *gin.Context) {
	middlewares.CreateOne(c, &models.Template{})
}

// 更新页面
func UpdateTemplateApi(c *gin.Context) {
	middlewares.UpdateOne(c, &models.Template{})
}

func DeleteTemplateApi(c *gin.Context) {
	middlewares.DeleteOne(c, &models.Template{})
}

func DeleteTemplateListApi(c *gin.Context) {
	middlewares.DeleteList(c, &models.Template{})
}
