package api_v2

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"goserver/libs/gorm"
	"goserver/middlewares"
	models "goserver/models/v2"
)

// 查询列表参数
type GetPublishListParams struct {
	ID      string `form:"id"`
	Title   string `form:"title"`
	Path    string `form:"path"`
	Version string `form:"version"`
	Online  string `form:"online"`
}

// 查询列表返回
type GetPublishListResponse []struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Path    string `json:"path"`
	Version string `json:"version"`
	Online  string `json:"online"`
}

func GetPublishListApi(c *gin.Context) {
	middlewares.GetList(c, &GetPublishListParams{}, &GetPublishListResponse{}, &models.Publish{})
}

// 查询版本参数
type GetPublishParams struct {
	Path string `form:"path"`
}

// 查询版本返回
type GetPublishResponse struct {
	models.Publish
}

func GetPublishApi(c *gin.Context) {
	middlewares.GetOne(c, &GetPublishParams{}, &GetPublishResponse{}, &models.Publish{})
}

// 版本发布
type PublishPublishParams struct {
	ID      string `json:"id"`
	Version string `json:"version"` // 发布的版本可回退，生成版本备份
	Path    string `json:"path"`    // 发布的版本可回退，生成版本备份
	Remark  string `json:"remark"`
}

// 新增版本
func CreatePublishApi(c *gin.Context) {
	var params PublishPublishParams
	middlewares.BindJSON(c, &params)
	page := models.Page{Base: models.Base{ID: params.ID}}
	gorm.First(&page)
	page.ID = ""
	page.CreateTime = nil
	page.UpdateTime = nil
	page.DeleteTime = nil
	model := models.Publish{Page: page, Version: params.Version, Remark: params.Remark}
	if err := gorm.Create(&model).Debug().Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 500, "message": "发布失败！", "err": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "发布成功！"})
	}
}

// 更新版本
func UpdatePublishApi(c *gin.Context) {
	middlewares.UpdateOne(c, &models.Publish{})
}

func DeletePublishApi(c *gin.Context) {
	middlewares.DeleteOne(c, &models.Publish{})
}
