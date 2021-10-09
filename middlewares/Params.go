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

func BindUri(c *gin.Context, params interface{}) error {
	if err := c.ShouldBindUri(params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}
	return nil
}

func BindQuery(c *gin.Context, params interface{}) error {
	if err := c.ShouldBindQuery(params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}
	return nil
}

func BindJSON(c *gin.Context, params interface{}) error {
	if err := c.ShouldBindJSON(params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}
	return nil
}

func GetTotal(c *gin.Context, db *gorm.DB) (int64, error) {
	var total int64
	if err := db.Count(&total).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "查询数据异常",
		})
		return total, err
	}
	return total, nil
}
