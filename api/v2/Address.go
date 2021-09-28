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
	var params models.Address
	if middlewares.ParseJSON(c, &params) {
		var data models.Address
		gorm.Where(&params).First(&data)
		if len(data.ID) > 0 {
			c.JSON(http.StatusOK, gin.H{"data": data})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "数据为空"})
	}
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
func GetAddressListApi(c *gin.Context) {
	middlewares.GetList(c, &AddressList{})
}

func CreateAddressListApi(c *gin.Context) {
	var params AddressList
	if middlewares.ParseJSON(c, &params) {
		for _, v := range params {
			gorm.Create(&v)
		}
		c.JSON(http.StatusOK, gin.H{"message": "创建成功！"})
	}
}

func UpdateAddressListApi(c *gin.Context) {
	var params AddressList
	if middlewares.ParseJSON(c, &params) {
		for _, v := range params {
			gorm.Updates(&models.Address{}, &v)
		}
		c.JSON(http.StatusOK, gin.H{"message": "更新成功！"})
	}
}

type DeleteListParams struct {
	IDs []string `json:"ids"`
}
func DeleteAddressListApi(c *gin.Context) {
	var params DeleteListParams
	if middlewares.ParseJSON(c, &params) {
		for _, v := range params.IDs {
			gorm.Delete(&models.Address{}, "id = ?", v)
		}
		c.JSON(http.StatusOK, gin.H{"message": "删除成功！"})	
	}
}
