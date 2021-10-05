package api_v2

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"goserver/libs/gorm"
	"goserver/middlewares"
	models "goserver/models/v2"
)

type AddressPlanList []models.AddressPlan

//
func GetAllAddressPlanApi(c *gin.Context) {
	middlewares.GetAll(c, &AddressPlanList{})
}

func GetAddressPlanApi(c *gin.Context) {
	middlewares.GetOne(c, &models.AddressPlan{}, &models.AddressPlan{})
}

func CreateAddressPlanApi(c *gin.Context) {
	middlewares.CreateOne(c, &models.AddressPlan{})
}

func UpdateAddressPlanApi(c *gin.Context) {
	middlewares.UpdateOne(c, &models.AddressPlan{}, &models.AddressPlan{})
}

func DeleteAddressPlanApi(c *gin.Context) {
	middlewares.DeleteOne(c, &models.AddressPlan{})
}

func GetAddressPlanListApi(c *gin.Context) {
	middlewares.GetList(c, &models.AddressPlan{}, &AddressPlanList{})
}

func CreateAddressPlanListApi(c *gin.Context) {
	var params AddressPlanList
	if middlewares.ParseParams(c, &params) {
		for _, v := range params {
			gorm.Create(&v)
		}
		c.JSON(http.StatusOK, gin.H{"message": "创建成功！"})
	}
}

func UpdateAddressPlanListApi(c *gin.Context) {
	var params AddressPlanList
	if middlewares.ParseParams(c, &params) {
		for _, v := range params {
			gorm.Updates(&models.AddressPlan{}, &v)
		}
		c.JSON(http.StatusOK, gin.H{"message": "更新成功！"})
	}
}

func DeleteAddressPlanListApi(c *gin.Context) {
	var params middlewares.DeleteListParams
	if middlewares.ParseParams(c, &params) {
		for _, v := range params.IDs {
			gorm.Delete(&models.AddressPlan{}, "id = ?", v)
		}
		c.JSON(http.StatusOK, gin.H{"message": "删除成功！"})
	}
}
