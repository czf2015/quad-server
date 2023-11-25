package api_v3

import (
	"time"

	"github.com/gin-gonic/gin"

	"goserver/middlewares"
	models "goserver/models/v3"
)

// 查询列表参数
type GetEventRuleListParams struct {
	Name                  string `gorm:"TYPE:varchar(255)" form:"name"`
	PrimaryClassification string `gorm:"TYPE:varchar(255)" form:"primary_classification"`
	ClassType             string `gorm:"TYPE:varchar(255)" form:"class_type"`
	ThreatLevel           string `gorm:"TYPE:varchar(255)" form:"threat_level"`
	Enable                bool   `gorm:"default:true" form:"enable"`
}

// 查询列表返回
type GetEventRuleListResponse []struct {
	ID                    uint       `json:"id"`
	Name                  string     `gorm:"TYPE:varchar(255)" json:"name"`
	PrimaryClassification string     `gorm:"TYPE:varchar(255)" json:"primary_classification"`
	ClassType             string     `gorm:"TYPE:varchar(255)" json:"class_type"`
	ThreatLevel           string     `gorm:"TYPE:varchar(255)" json:"threat_level"`
	Enable                bool       `gorm:"default:true" json:"enable"`
	CreatedAt             *time.Time `gorm:"type:timestamp;default:NOW()" json:"create_time"`
	UpdatedAt             *time.Time `gorm:"type:timestamp;default:NOW()" json:"update_time"`
}

func GetEventRuleListApi(c *gin.Context) {
	middlewares.GetList(c, &GetEventRuleListParams{}, &GetEventRuleListResponse{}, &models.EventRule{})
}

// 查询页面参数
type GetEventRuleParams struct {
	ID uint `form:"id"`
}

// 查询页面返回
type GetEventRuleResponse struct {
	models.EventRule
}

func GetEventRuleApi(c *gin.Context) {
	middlewares.GetOne(c, &GetEventRuleParams{}, &GetEventRuleResponse{}, &models.EventRule{})
}

// 新增页面
func CreateEventRuleApi(c *gin.Context) {
	middlewares.CreateOne(c, &models.EventRule{})
}

// 更新页面
func UpdateEventRuleApi(c *gin.Context) {
	middlewares.UpdateOne(c, &models.EventRule{})
}

func DeleteEventRuleApi(c *gin.Context) {
	middlewares.DeleteOne(c, &models.EventRule{})
}

func DeleteEventRuleListApi(c *gin.Context) {
	middlewares.DeleteList(c, &models.EventRule{})
}
