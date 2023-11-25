package api_v3

import (
	models "goserver/models/v3"

	"goserver/middlewares"

	"github.com/gin-gonic/gin"
)

func CreatePermission(c *gin.Context) {
	middlewares.Create(c, &models.Permission{})
}

func GetPermission(c *gin.Context) {
	middlewares.Get(c, &models.Permission{})
}

func UpdatePermission(c *gin.Context) {
	middlewares.Update(c, &models.Permission{})
}

func DeletePermission(c *gin.Context) {
	middlewares.Delete(c, &models.Permission{})
}
