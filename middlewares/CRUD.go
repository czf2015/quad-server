package middlewares

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"goserver/libs/gorm"
)

func GetAll(c *gin.Context, data interface{}) {
	gorm.Find(data)
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": &data})
}

func GetOne(c *gin.Context, params, data, model interface{}) {
	if BindQuery(c, params) {
		db := gorm.GetDB().Model(model).Where(params)
		if total, ok := GetTotal(c, db); ok {
			if total > 0 {
				db.First(data)
				c.JSON(http.StatusOK, gin.H{"code": 200, "message": "查询成功！", "data": data})
				return
			}
			c.JSON(http.StatusOK, gin.H{"code": 500, "message": "查询不到！"})
		}
	}
}

func GetList(c *gin.Context, params, data, model interface{}) {
	if BindQuery(c, params) {
		fmt.Println(params)
		db := gorm.GetDB().Model(model).Where(params).Debug()
		if total, ok := GetTotal(c, db); ok {
			offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
			limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
			order := c.DefaultQuery("order", "desc")
			if err := db.Order("update_time " + order).Limit(limit).Offset(offset).Find(data).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code":    500,
					"message": "查询数据异常",
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "success",
				"data": map[string]interface{}{
					"list":     data,
					"total":    total,
				},
			})
		}
	}
}

func CreateOne(c *gin.Context, model interface{}) {
	if BindJSON(c, model) {
		if err := gorm.Create(model).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": 500, "message": "创建失败！", "err": err})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 200, "message": "创建成功！", "data": model})
		}
	}
}

// TOFIX:
func CreateList(c *gin.Context, params, model interface{}) {
	if BindJSON(c, params) {
		gorm.Create(model)
		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "创建成功！"})
	}
}

func UpdateOne(c *gin.Context, model interface{}) {
	if BindJSON(c, model) {
		if err := gorm.Updates(model, c.PostForm("id")).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "message": "更新失败！", "err": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "更新成功！", "data": model})
	}
}

func DeleteOne(c *gin.Context, model interface{}) {
	gorm.DeleteByID(model, c.Query("id")).Debug()
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "删除成功！"})
}

func DeleteList(c *gin.Context, model interface{}) {
	var params DeleteListParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "参数错误！", "err": err})
		return
	}

	// 在数据库中删除符合条件的记录
	result := gorm.Delete(model, "id in ?", params.IDs)
	if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "删除失败！", "error": result.Error.Error()})
			return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
			"message": "删除成功！",
			"code": 200,
	})
}
