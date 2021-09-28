package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ParseJSON(c *gin.Context, params interface{}) bool {
	if err := c.ShouldBindJSON(params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}
	return true
}