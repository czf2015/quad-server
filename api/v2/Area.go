package api_v2

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"goserver/libs/gorm"
	"goserver/middlewares"
	models "goserver/models/v2"
)

type AreaList []models.Area

func GetAllAreaApi(c *gin.Context) {
	middlewares.GetAll(c, &AreaList{})
}

func GetAreaApi(c *gin.Context) {
	middlewares.GetOne(c, &models.Area{}, &models.Area{}, &models.Area{})
}

func GetAreaListApi(c *gin.Context) {
	middlewares.GetList(c, &models.Area{}, &AreaList{}, &models.Area{})
}

func CreateAreaApi(c *gin.Context) {
	middlewares.CreateOne(c, &models.Area{})
}

func CreateAreaListApi(c *gin.Context) {
	middlewares.CreateList(c, &AreaList{}, &models.Area{})
}

func UpdateAreaApi(c *gin.Context) {
	middlewares.UpdateOne(c, &models.Area{})
}

func UpdateAreaListApi(c *gin.Context) {
	var params AreaList
	if middlewares.BindJSON(c, &params) {
		for _, v := range params {
			gorm.Updates(&v)
		}
		c.JSON(http.StatusOK, gin.H{"message": "更新成功！"})
	}
}

func DeleteAreaApi(c *gin.Context) {
	middlewares.DeleteOne(c, &models.Area{})
}

func DeleteAreaListApi(c *gin.Context) {
	middlewares.DeleteList(c, &models.Area{})
}
