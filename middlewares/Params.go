package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteParams struct {
	ID string `json:"id"`
}

type DeleteListParams struct {
	IDs []string `json:"ids"`
}

func ParseParams(c *gin.Context, params interface{}) bool {
	if err := c.ShouldBindJSON(params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}
	return true
}
