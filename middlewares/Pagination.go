package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Page     int
	PageSize int
}

func GetPagination(c *gin.Context) (Pagination, error) {
	var params Pagination
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return params, err
	}
	return params, nil
}
