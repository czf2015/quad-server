package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"goserver/libs/gorm"
)

func GetList(c *gin.Context, data interface{}) {
	gorm.Find(data)
	c.JSON(http.StatusOK, gin.H{"data": &data})
}

func CreateOne(c *gin.Context, params interface{}) {
	if ParseJSON(c, params) {
		gorm.Create(params)
		c.JSON(http.StatusOK, gin.H{"message": "创建成功！"})
	}
}


func UpdateOne(c *gin.Context, model, params interface{}) {
	if ParseJSON(c, params) {
		gorm.Updates(model, params)
		c.JSON(http.StatusOK, gin.H{"message": "更新成功！"})
	}
}

type DeleteParams struct {
	ID string `json:"id"`
}
func DeleteOne(c *gin.Context, model interface{}) {
	var params DeleteParams
	if ParseJSON(c, &params) {
		gorm.Delete(model, "id = ?", params.ID)
		c.JSON(http.StatusOK, gin.H{"message": "删除成功！"})	
	}
}
