package api_v2

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"goserver/libs/gorm"
	"goserver/middlewares"
	models "goserver/models/v2"
)


type AreaList []models.Area

func GetAreaApi(c *gin.Context) {
	var params models.Area
	if middlewares.ParseParams(c, &params) {
		var data models.Area
		gorm.Where(&params).First(&data)
		if len(data.ID) > 0 {
			c.JSON(http.StatusOK, gin.H{"data": data})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "数据为空"})
	}
}

func CreateAreaApi(c *gin.Context) {
	middlewares.CreateOne(c, &models.Area{})
}

func UpdateAreaApi(c *gin.Context) {
	middlewares.UpdateOne(c, &models.Area{}, &models.Area{})
}

func DeleteAreaApi(c *gin.Context) {	
	middlewares.DeleteOne(c, &models.Area{})
}

// 
func GetAllAreaApi(c *gin.Context) {
	middlewares.GetAll(c, &AreaList{})
}

func CreateAreaListApi(c *gin.Context) {
	var params AreaList
	if middlewares.ParseParams(c, &params) {
		for _, v := range params {
			gorm.Create(&v)
		}
		c.JSON(http.StatusOK, gin.H{"message": "创建成功！"})
	}
}

func UpdateAreaListApi(c *gin.Context) {
	var params AreaList
	if middlewares.ParseParams(c, &params) {
		for _, v := range params {
			gorm.Updates(&models.Area{}, &v)
		}
		c.JSON(http.StatusOK, gin.H{"message": "更新成功！"})
	}
}

func DeleteAreaListApi(c *gin.Context) {
	var params middlewares.DeleteListParams
	if middlewares.ParseParams(c, &params) {
		for _, v := range params.IDs {
			gorm.Delete(&models.Area{}, "id = ?", v)
		}
		c.JSON(http.StatusOK, gin.H{"message": "删除成功！"})	
	}
}
