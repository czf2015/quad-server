package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"goserver/libs/gorm"
)

type DeleteParams struct {
	ID string `json:"id"`
}

type DeleteListParams struct {
	IDs []string `json:"ids"`
}

func BindUri(c *gin.Context, params interface{}) bool {
	ok := true
	if err := c.ShouldBindUri(params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ok = false
	}
	return ok
}

func BindQuery(c *gin.Context, params interface{}) bool {
	ok := true
	if err := c.ShouldBindQuery(params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ok = false
	}
	return ok
}

func BindJSON(c *gin.Context, params interface{}) bool {
	ok := true
	if err := c.ShouldBindJSON(params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ok = false
	}
	return ok
}

func GetTotal(c *gin.Context, db *gorm.DB) (int64, bool) {
	var total int64
	ok := true
	if err := db.Count(&total).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "查询数据异常",
		})
		ok = false
	}
	return total, ok
}
