// undefined
package api_v2

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"goserver/libs/gorm"
	"goserver/middlewares"
	models "goserver/models/v2"
)

type NetworkManageList []models.NetworkManage

func GetAllNetworkManageApi(c *gin.Context) {
	middlewares.GetAll(c, &NetworkManageList{})
}

func GetNetworkManageApi(c *gin.Context) {
	middlewares.GetOne(c, &models.NetworkManage{}, &models.NetworkManage{})
}

func GetNetworkManageListApi(c *gin.Context) {
	middlewares.GetList(c, &models.NetworkManage{}, &NetworkManageList{})
}

func CreateNetworkManageApi(c *gin.Context) {
	middlewares.CreateOne(c, &models.NetworkManage{})
}

func CreateNetworkManageListApi(c *gin.Context) {
	middlewares.CreateList(c, &NetworkManageList{})
}

func UpdateNetworkManageApi(c *gin.Context) {
	middlewares.UpdateOne(c, &models.NetworkManage{})
}

func UpdateNetworkManageListApi(c *gin.Context) {
	var params NetworkManageList
	if middlewares.BindJSON(c, &params) {
		for _, v := range params {
			gorm.Updates(&v)
		}
		c.JSON(http.StatusOK, gin.H{"message": "更新成功！"})
	}
}

func DeleteNetworkManageApi(c *gin.Context) {
	middlewares.DeleteOne(c, &models.NetworkManage{})
}

func DeleteNetworkManageListApi(c *gin.Context) {
	middlewares.DeleteList(c, &models.NetworkManage{})
}
