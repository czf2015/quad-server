package api_v2

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"goserver/libs/gorm"
	"goserver/middlewares"
	models "goserver/models/v2"
)

type AddressList []models.Address

func GetAllAddressApi(c *gin.Context) {
	middlewares.GetAll(c, &AddressList{})
}

func GetAddressApi(c *gin.Context) {
	middlewares.GetOne(c, &models.Address{}, &models.Address{})
}

func GetAddressListApi(c *gin.Context) {
	middlewares.GetList(c, &models.Address{}, &AddressList{})
}

func CreateAddressApi(c *gin.Context) {
	middlewares.CreateOne(c, &models.Address{})
}

func CreateAddressListApi(c *gin.Context) {
	middlewares.CreateList(c, &AddressList{})
}

func UpdateAddressApi(c *gin.Context) {
	middlewares.UpdateOne(c, &models.Address{})
}

func UpdateAddressListApi(c *gin.Context) {
	var params AddressList
	if middlewares.BindJSON(c, &params) == nil {
		for _, v := range params {
			gorm.Updates(&v)
		}
		c.JSON(http.StatusOK, gin.H{"message": "更新成功！"})
	}
}

func DeleteAddressApi(c *gin.Context) {
	middlewares.DeleteOne(c, &models.Address{})
}

func DeleteAddressListApi(c *gin.Context) {
	middlewares.DeleteList(c, &models.Address{})
}
