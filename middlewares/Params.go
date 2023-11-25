package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func BindUri(c *gin.Context, params interface{}) bool {
	ok := true
	if err := c.ShouldBindUri(params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		ok = false
	}
	return ok
}

func BindQuery(c *gin.Context, params interface{}) bool {
	ok := true
	if err := c.ShouldBindQuery(params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		ok = false
	}
	return ok
}

func BindJSON(c *gin.Context, params interface{}) bool {
	ok := true
	if err := c.ShouldBindJSON(params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		ok = false
	}
	return ok
}

func GetTotal(c *gin.Context, db *gorm.DB) (int64, bool) {
	var total int64
	ok := true
	if err := db.Count(&total).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "查询数据异常",
		})
		ok = false
	}
	return total, ok
}

// 分页封装
func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
