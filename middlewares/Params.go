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

func BindJSON(c *gin.Context, params interface{}) error {
	if err := c.ShouldBindJSON(params); err != nil {
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
