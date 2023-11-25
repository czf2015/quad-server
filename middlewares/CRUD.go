package middlewares

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"goserver/libs/e"
)

func GetAll(c *gin.Context, data interface{}) {
	code := e.SUCCESS
	if err := db.Find(data).Error; err != nil {
		code = e.ERROR_ORM_GET
		c.JSON(http.StatusInternalServerError, gin.H{"code": code, "message": e.GetMsg(code)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": code, "message": "查询成功！", "data": &data})
}

func GetOne(c *gin.Context, params, data, model interface{}) {
	code := e.SUCCESS
	if BindQuery(c, params) {
		_db := db.Model(model).Where(params)
		if total, ok := GetTotal(c, _db); ok {
			if total > 0 {
				if err := _db.First(data).Error; err != nil {
					code = e.ERROR_ORM_GET
					c.JSON(http.StatusInternalServerError, gin.H{"code": code, "message": e.GetMsg(code)})
					return
				}
				c.JSON(http.StatusOK, gin.H{"code": code, "message": "查询成功！", "data": data})
			} else {
				code = e.ERROR_NOT_FOUND
				c.JSON(http.StatusNotFound, gin.H{"code": code, "message": e.GetMsg(code)})
			}
		}
	}
}

func GetList(c *gin.Context, params, data, model interface{}) {
	code := e.SUCCESS
	if BindQuery(c, params) {
		_db := db.Model(model).Where(params).Debug()
		if total, ok := GetTotal(c, _db); ok {
			offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
			limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
			order := c.DefaultQuery("order", "desc")
			if err := _db.Order("created_at " + order).Limit(limit).Offset(offset).Find(data).Debug().Error; err != nil {
				code = e.ERROR_ORM_GET
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    code,
					"message": "查询数据异常",
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": e.GetMsg(code),
				"data": map[string]interface{}{
					"list":  data,
					"total": total,
				},
			})
		}
	}
}

func CreateOne(c *gin.Context, model interface{}) {
	code := e.SUCCESS
	if BindJSON(c, model) {
		if err := db.Create(model).Error; err != nil {
			code = e.ERROR_ORM_CREATE
			c.JSON(http.StatusInternalServerError, gin.H{"code": code, "message": e.GetMsg(code), "err": err})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": code, "message": "创建成功！", "data": model})
		}
	}
}

// TOFIX:
func CreateList(c *gin.Context, params, model interface{}) {
	if BindJSON(c, params) {
		db.Create(model)
		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "创建成功！"})
	}
}

func UpdateOne(c *gin.Context, model interface{}) {
	code := e.SUCCESS
	_db := db.Model(model).Where("id = ?", c.Param("id"))
	if total, ok := GetTotal(c, _db); ok {
		if total > 0 {

			// c.JSON(http.StatusOK, gin.H{"code": code, "message": "查询成功！", "data": data})
			if BindJSON(c, model) {
				if err := _db.Updates(model).Debug().Error; err != nil {
					code = e.ERROR_ORM_UPDATE
					c.JSON(http.StatusInternalServerError, gin.H{"code": code, "message": e.GetMsg(code), "err": err})
					return
				}
				c.JSON(http.StatusOK, gin.H{"code": code, "message": "更新成功！", "data": model})
			}
		} else {
			code = e.ERROR_NOT_FOUND
			c.JSON(http.StatusNotFound, gin.H{"code": code, "message": e.GetMsg(code)})
		}
	}

}

func DeleteOne(c *gin.Context, model interface{}) {
	code := e.SUCCESS
	if err := db.Delete(model, c.Param("id")).Error; err != nil {
		code = e.ERROR_ORM_DELETE
		c.JSON(http.StatusInternalServerError, gin.H{"code": code, "message": e.GetMsg(code), "err": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": code, "message": "删除成功！"})
}

type DeleteListParams struct {
	IDs []interface{} `json:"ids"`
}

func DeleteList(c *gin.Context, model interface{}) {
	code := e.SUCCESS
	var params DeleteListParams
	if err := c.ShouldBindJSON(&params); err != nil {
		code = e.INVALID_PARAMS
		c.JSON(http.StatusBadRequest, gin.H{"code": code, "message": e.GetMsg(code), "err": err})
		return
	}

	if err := db.Delete(model, params.IDs).Error; err != nil {
		code = e.ERROR_ORM_DELETE
		c.JSON(http.StatusInternalServerError, gin.H{"code": code, "message": e.GetMsg(code), "err": err})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": "删除成功！",
	})
}
