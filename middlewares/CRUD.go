package middlewares

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"goserver/libs/gorm"
)

func GetTotal(c *gin.Context, db *gorm.DB) (int, error) {
	var total int
	if err := db.Count(&total).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "查询数据异常",
		})
		return total, err
	}
	return total, nil
}

func GetAll(c *gin.Context, data interface{}) {
	gorm.Find(data)
	c.JSON(http.StatusOK, gin.H{"data": &data})
}

func GetList(c *gin.Context, model, data interface{}) {
	if BindQuery(c, model) == nil {
		db := gorm.GetDB().Model(model).Where(model)
		if total, err := GetTotal(c, db); err == nil {
			page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
			pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
			offset := (page - 1) * pageSize
			if err := db.Order("id desc")/* .Where(model) */.Limit(pageSize).Offset(offset).Find(data).Error; err != nil {
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

func GetOne(c *gin.Context, params, data interface{}) {
	if BindJSON(c, params) == nil {
		db := gorm.GetDB().Model(params).Where(params)

		var total int
		if err := db.Count(&total).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    500,
				"message": "查询数据异常",
			})
			return
		}
		if total > 0 {
			db.First(data)
			c.JSON(http.StatusOK, gin.H{"data": data})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "数据为空"})
	}
}

func CreateOne(c *gin.Context, params interface{}) {
	if BindJSON(c, params) == nil {
		gorm.Create(params)
		c.JSON(http.StatusOK, gin.H{"message": "创建成功！"})
	}
}

func UpdateOne(c *gin.Context, model, params interface{}) {
	if BindJSON(c, params) == nil {
		gorm.Updates(model, params)
		c.JSON(http.StatusOK, gin.H{"message": "更新成功！"})
	}
}

func DeleteOne(c *gin.Context, model interface{}) {
	var params DeleteParams
	if BindJSON(c, &params) == nil {
		gorm.Delete(model, "id = ?", params.ID)
		c.JSON(http.StatusOK, gin.H{"message": "删除成功！"})
	}
}
