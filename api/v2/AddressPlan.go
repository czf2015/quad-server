package api_v2

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"goserver/libs/gorm"
	"goserver/middlewares"
	models "goserver/models/v2"
)

type AddressPlanList []models.AddressPlan

func GetAllAddressPlanApi(c *gin.Context) {
	middlewares.GetAll(c, &AddressPlanList{})
}

func GetAddressPlanApi(c *gin.Context) {
	middlewares.GetOne(c, &models.AddressPlan{}, &models.AddressPlan{})
}

func GetAddressPlanListApi(c *gin.Context) {
	middlewares.GetList(c, &models.AddressPlan{}, &AddressPlanList{})
}

func CreateAddressPlanApi(c *gin.Context) {
	middlewares.CreateOne(c, &models.AddressPlan{})
}

func CreateAddressPlanListApi(c *gin.Context) {
	middlewares.CreateList(c, &AddressPlanList{})
}

func UpdateAddressPlanApi(c *gin.Context) {
	middlewares.UpdateOne(c, &models.AddressPlan{})
}

func UpdateAddressPlanListApi(c *gin.Context) {
	var params AddressPlanList
	if middlewares.BindJSON(c, &params) == nil {
		for _, v := range params {
			gorm.Updates(&v)
		}
		c.JSON(http.StatusOK, gin.H{"message": "更新成功！"})
	}
}

func DeleteAddressPlanApi(c *gin.Context) {
	middlewares.DeleteOne(c, &models.AddressPlan{})
}

func DeleteAddressPlanListApi(c *gin.Context) {
	middlewares.DeleteList(c, &models.AddressPlan{})
}
