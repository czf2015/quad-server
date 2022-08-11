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
				c.JSON(http.StatusOK, gin.H{"code": 200, "message": "创建成功！", "data": data})
				return
			}
			c.JSON(http.StatusOK, gin.H{"code": 500, "message": "数据为空"})
		}
	}
}

func GetList(c *gin.Context, params, data, model interface{}) {
	if BindQuery(c, params) {
		fmt.Println(params)
		db := gorm.GetDB().Model(model).Where(params).Debug()
		if total, ok := GetTotal(c, db); ok {
			page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
			pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
			offset := (page - 1) * pageSize
			if err := db.Order("id desc").Limit(pageSize).Offset(offset).Find(data).Error; err != nil {
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
					"page":     page,
					"pageSize": pageSize,
				},
			})
			return
		}
	}
}

func CreateOne(c *gin.Context, model interface{}) {
	if BindJSON(c, model) {
		if err := gorm.Create(model).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": 500, "message": "操作失败！", "err": err})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 200, "message": "操作成功！", "data": model})
		}
	}
}

func CreateList(c *gin.Context, params, model interface{}) {
	if BindJSON(c, params) {
		gorm.Create(model)
		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "创建成功！"})
	}
}

func UpdateOne(c *gin.Context, model interface{}) {
	if BindJSON(c, model) {
		if err := gorm.Updates(model); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "message": "更新失败！", "err": err})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 200, "message": "更新成功！", "data": model})
		}
	}
}

func DeleteOne(c *gin.Context, model interface{}) {
	var params DeleteParams
	if BindQuery(c, &params) {
		gorm.DeleteByID(model, params.ID).Debug()
		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "删除成功！"})
	}
}

func DeleteList(c *gin.Context, model interface{}) {
	var params DeleteListParams
	if BindJSON(c, &params) {
		gorm.DeleteByID(model, params.IDs)
		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "删除成功！"})
	}
}
