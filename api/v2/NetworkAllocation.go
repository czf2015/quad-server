// undefined
package api_v2

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"goserver/libs/gorm"
	"goserver/middlewares"
	models "goserver/models/v2"
)

type NetworkAllocationList []models.NetworkAllocation

func GetAllNetworkAllocationApi(c *gin.Context) {
	middlewares.GetAll(c, &NetworkAllocationList{})
}

func GetNetworkAllocationApi(c *gin.Context) {
	middlewares.GetOne(c, &models.NetworkAllocation{}, &models.NetworkAllocation{}, &models.NetworkAllocation{})
}

func GetNetworkAllocationListApi(c *gin.Context) {
	middlewares.GetList(c, &models.NetworkAllocation{}, &NetworkAllocationList{}, &models.NetworkAllocation{})
}

func CreateNetworkAllocationApi(c *gin.Context) {
	middlewares.CreateOne(c, &models.NetworkAllocation{})
}

func CreateNetworkAllocationListApi(c *gin.Context) {
	middlewares.CreateList(c, &NetworkAllocationList{}, &models.NetworkAllocation{})
}

func UpdateNetworkAllocationApi(c *gin.Context) {
	middlewares.UpdateOne(c, &models.NetworkAllocation{})
}

func UpdateNetworkAllocationListApi(c *gin.Context) {
	var params NetworkAllocationList
	if middlewares.BindJSON(c, &params) {
		for _, v := range params {
			gorm.Updates(&v)
		}
		c.JSON(http.StatusOK, gin.H{"message": "更新成功！"})
	}
}

func DeleteNetworkAllocationApi(c *gin.Context) {
	middlewares.DeleteOne(c, &models.NetworkAllocation{})
}

func DeleteNetworkAllocationListApi(c *gin.Context) {
	middlewares.DeleteList(c, &models.NetworkAllocation{})
}
