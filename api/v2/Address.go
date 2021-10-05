package api_v2

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"goserver/libs/gorm"
	"goserver/middlewares"
	models "goserver/models/v2"
)

type AddressList []models.Address

func GetAddressApi(c *gin.Context) {
	middlewares.GetOne(c, &models.Address{}, &models.Address{})
}

func CreateAddressApi(c *gin.Context) {
	middlewares.CreateOne(c, &models.Address{})
}

func UpdateAddressApi(c *gin.Context) {
	middlewares.UpdateOne(c, &models.Address{}, &models.Address{})
}

func DeleteAddressApi(c *gin.Context) {
	middlewares.DeleteOne(c, &models.Address{})
}

//
func GetAllAddressApi(c *gin.Context) {
	middlewares.GetAll(c, &AddressList{})
}

func CreateAddressListApi(c *gin.Context) {
	var params AddressList
	if middlewares.ParseParams(c, &params) {
		for _, v := range params {
			gorm.Create(&v)
		}
		c.JSON(http.StatusOK, gin.H{"message": "创建成功！"})
	}
}

func UpdateAddressListApi(c *gin.Context) {
	var params AddressList
	if middlewares.ParseParams(c, &params) {
		for _, v := range params {
			gorm.Updates(&models.Address{}, &v)
		}
		c.JSON(http.StatusOK, gin.H{"message": "更新成功！"})
	}
}

func DeleteAddressListApi(c *gin.Context) {
	var params middlewares.DeleteListParams
	if middlewares.ParseParams(c, &params) {
		for _, v := range params.IDs {
			gorm.Delete(&models.Address{}, "id = ?", v)
		}
		c.JSON(http.StatusOK, gin.H{"message": "删除成功！"})
	}
}
