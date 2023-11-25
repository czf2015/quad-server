package api_v3

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"goserver/middlewares"
	models "goserver/models/v3"
)

// 查询列表参数
type GetPublishListParams struct {
	ID      uint   `form:"id"`
	Title   string `form:"title"`
	Path    string `form:"path"`
	Version string `form:"version"`
	Online  int    `form:"online"`
}

// 查询列表返回
type GetPublishListResponse []struct {
	ID        uint       `json:"id"`
	Title     string     `json:"title"`
	Path      string     `json:"path"`
	Version   string     `json:"version"`
	Online    int        `json:"online"`
	ImgUrl    string     `gorm:"type:varchar(255);not null;column:img_url" json:"imgUrl"`
	CreatedAt *time.Time `json:"create_time"`
}

func GetPublishListApi(c *gin.Context) {
	middlewares.GetList(c, &GetPublishListParams{}, &GetPublishListResponse{}, &models.Publish{})
}

// 查询版本参数
type GetPublishParams struct {
	ID     uint   `form:"id"`
	Path   string `form:"path"`
	Online int    `form:"online"`
}

// 查询版本返回
type GetPublishResponse struct {
	models.Publish
}

func GetPublishApi(c *gin.Context) {
	middlewares.GetOne(c, &GetPublishParams{}, &GetPublishResponse{}, &models.Publish{})
}

// 版本发布
type CreatePublishParams struct {
	ID      uint   `json:"id"`
	Version string `json:"version"` // 发布的版本可回退，生成版本备份
	Path    string `json:"path"`
	Remark  string `json:"remark"`
}

// 新增版本
func CreatePublishApi(c *gin.Context) {
	var params CreatePublishParams
	if middlewares.BindJSON(c, &params) {
		page := models.Page{Base: models.Base{ID: params.ID}}
		db.First(&page)
		page.ID = 0
		page.CreatedAt = nil
		page.UpdatedAt = nil
		page.DeletedAt = nil
		page.Path = params.Path
		model := models.Publish{Page: page, Version: params.Version, Remark: params.Remark}
		if err := db.Create(&model).Debug().Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "发布失败！", "err": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "发布成功！"})
	}
}

// 更新版本
func UpdatePublishApi(c *gin.Context) {
	middlewares.UpdateOne(c, &models.Publish{})
}

type PatchPublishParams struct {
	Online int `json:"online"`
}
type PatchPublishResponse struct {
	ID      uint   `json:"id"`
	Title   string `form:"title"`
	Path    string `form:"path"`
	Version string `form:"version"`
	Online  int    `json:"online"`
}

func PatchPublishApi(c *gin.Context) {
	var params PatchPublishParams
	var data PatchPublishResponse
	if middlewares.BindJSON(c, &params) {
		db := db.Model(&models.Publish{}).Where("id = ?", c.Param("id"))
		if total, ok := middlewares.GetTotal(c, db); ok {
			if total > 0 {
				if err := db.Update("Online", params.Online).Error; err != nil {
					// 处理更新错误
					c.JSON(http.StatusOK, gin.H{"code": 500, "message": "更新失败！", "err": err})
					return
				}

				db.First(&data)
				c.JSON(http.StatusOK, gin.H{"code": 200, "message": "更新成功！", "data": data})
				return
			}

			c.JSON(http.StatusOK, gin.H{"code": 500, "message": "查询不到！", "err": nil})
		}
	}
}

func DeletePublishApi(c *gin.Context) {
	middlewares.DeleteOne(c, &models.Publish{})
}

func DeletePublishListApi(c *gin.Context) {
	middlewares.DeleteList(c, &models.Publish{})
}
