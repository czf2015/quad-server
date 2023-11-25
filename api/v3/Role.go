package api_v3

import (
	models "goserver/models/v3"

	"goserver/middlewares"

	"github.com/gin-gonic/gin"
)

func CreateRole(c *gin.Context) {
	middlewares.Create(c, &models.Role{})
}

func GetRole(c *gin.Context) {
	middlewares.Get(c, &models.Role{})
}

func UpdateRole(c *gin.Context) {
	middlewares.Update(c, &models.Role{})
}

func DeleteRole(c *gin.Context) {
	middlewares.Delete(c, &models.Role{})
}
