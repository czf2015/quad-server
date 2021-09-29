package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"goserver/libs/gorm"
)

func GetAll(c *gin.Context, data interface{}) {
	gorm.Find(data)
	c.JSON(http.StatusOK, gin.H{"data": &data})
}

func CreateOne(c *gin.Context, params interface{}) {
	if ParseParams(c, params) {
		gorm.Create(params)
		c.JSON(http.StatusOK, gin.H{"message": "创建成功！"})
	}
}


func UpdateOne(c *gin.Context, model, params interface{}) {
	if ParseParams(c, params) {
		gorm.Updates(model, params)
		c.JSON(http.StatusOK, gin.H{"message": "更新成功！"})
	}
}

func DeleteOne(c *gin.Context, model interface{}) {
	var params DeleteParams
	if ParseParams(c, &params) {
		gorm.Delete(model, "id = ?", params.ID)
		c.JSON(http.StatusOK, gin.H{"message": "删除成功！"})	
	}
}
