package api_v3

import (
	"time"

	"github.com/gin-gonic/gin"

	"goserver/middlewares"
	models "goserver/models/v3"
)

// 查询列表参数
type GetPageListParams struct {
	Title string `form:"title"`
	Path  string `form:"path"`
	// Tags  []string `form:"tags"`
}

// 查询列表返回
type GetPageListResponse []struct {
	ID        uint             `json:"id"`
	Title     string           `json:"title"`
	Path      string           `json:"path"`
	Tags      models.FlatArray `gorm:"TYPE:json" json:"tags"`
	CreatedAt *time.Time       `gorm:"type:timestamp;default:NOW()" json:"create_time"`
	ImgUrl    string           `gorm:"type:varchar(255);not null;column:img_url" json:"imgUrl"`
	UpdatedAt *time.Time       `gorm:"type:timestamp;default:NOW()" json:"update_time"`
}

func GetPageListApi(c *gin.Context) {
	middlewares.GetList(c, &GetPageListParams{}, &GetPageListResponse{}, &models.Page{})
}

// 查询页面参数
type GetPageParams struct {
	ID uint `form:"id"`
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

func DeletePageListApi(c *gin.Context) {
	middlewares.DeleteList(c, &models.Page{})
}