package api_v2

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"goserver/libs/gorm"
	models "goserver/models/v2"
)

type AddressList []models.Address

func GetAddressListApi(c *gin.Context) {
	var addressList AddressList
	gorm.Find(&addressList)
	c.JSON(http.StatusOK, gin.H{"data": addressList})
}

func GetAddressApi(c *gin.Context) {
	var params models.Address
	var data models.Address
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	gorm.Where(params).First(&data)
	if len(data.ID) > 0 {
		c.JSON(http.StatusOK, gin.H{"data": data})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "数据为空"})
}

func CreateAddressListApi(c *gin.Context) {
	var params AddressList
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, v := range params {
		gorm.Create(&v)
	}
	c.JSON(http.StatusOK, gin.H{"message": "创建成功！"})
}

func CreateAddressApi(c *gin.Context) {
	var params models.Address
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	gorm.Create(&params)
	c.JSON(http.StatusOK, gin.H{"message": "创建成功！"})
}

func UpdateAddressListApi(c *gin.Context) {
	var params AddressList
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, v := range params {
		gorm.Updates(&models.Address{}, &v)
	}
	c.JSON(http.StatusOK, gin.H{"message": "更新成功！"})
}

func UpdateAddressApi(c *gin.Context) {
	var params models.Address
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	gorm.Updates(&models.Address{}, &params)
	c.JSON(http.StatusOK, gin.H{"message": "更新成功！"})
}

type DeleteListParams struct {
	IDs []string `json:"ids"`
}
func DeleteAddressListApi(c *gin.Context) {
	var params DeleteListParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, v := range params.IDs {
		gorm.Delete(&models.Address{}, "id = ?", v)
	}
	c.JSON(http.StatusOK, gin.H{"message": "删除成功！"})	
}

type DeleteParams struct {
	ID string `json:"id"`
}
func DeleteAddressApi(c *gin.Context) {
	var params DeleteParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	gorm.Delete(&models.Address{}, "id = ?", params.ID)
	c.JSON(http.StatusOK, gin.H{"message": "删除成功！"})	
}