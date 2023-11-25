package middlewares

import (
	"net/http"

	"goserver/libs/e"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context, model interface{}) {
	code := e.SUCCESS
	if err := c.ShouldBindJSON(model); err != nil {
		code = e.INVALID_PARAMS
		c.JSON(http.StatusBadRequest, gin.H{"code": code, "message": e.GetMsg(code), "err": err.Error()})
		return
	}

	if err := db.Create(model).Error; err != nil {
		code = e.ERROR_ORM_CREATE
		c.JSON(http.StatusInternalServerError, gin.H{"code": code, "message": e.GetMsg(code), "err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": code, "message": "创建成功！", "data": model})
}

func Get(c *gin.Context, model interface{}) {
	code := e.SUCCESS
	if err := db.First(model, c.Param("id")).Error; err != nil {
		code = e.ERROR_ORM_GET
		c.JSON(http.StatusInternalServerError, gin.H{"code": code, "message": e.GetMsg(code), "err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": code, "message": "查询成功！", "data": model})
}

func Update(c *gin.Context, model interface{}) {
	code := e.SUCCESS

	if err := db.First(model, c.Param("id")).Error; err != nil {
		code = e.ERROR_ORM_GET
		c.JSON(http.StatusNotFound, gin.H{"code": code, "message": e.GetMsg(code), "err": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(model); err != nil {
		code = e.INVALID_PARAMS
		c.JSON(http.StatusBadRequest, gin.H{"code": code, "message": e.GetMsg(code), "err": err.Error()})
		return
	}

	if err := db.Save(model).Error; err != nil {
		code = e.ERROR_ORM_UPDATE
		c.JSON(http.StatusInternalServerError, gin.H{"code": code, "message": e.GetMsg(code), "err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": code, "message": "更新成功！", "data": model})
}

func Delete(c *gin.Context, model interface{}) {
	code := e.SUCCESS
	if err := db.Delete(model, c.Param("id")).Error; err != nil {
		code = e.ERROR_ORM_DELETE
		c.JSON(http.StatusInternalServerError, gin.H{"code": code, "message": e.GetMsg(code), "err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": "删除成功！",
	})
}
