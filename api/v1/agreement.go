package api_v1

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"goserver/libs/jwt"
	"goserver/models"
	"goserver/libs/gorm"
)

type SignedAgreementRequest struct {
	Id string `json:"id"`
}

func UnsignedAgreementsApi(c *gin.Context) {
	claims, _ := jwt.ParseToken(c.Query("token"))
	user := models.GetUserById(claims.Id)

	c.JSON(http.StatusOK, gin.H{
		"unsigned": user.GetLatestUnsignedAgreements(),
	})
}

func SignAgreementApi(c *gin.Context) {
	claims, _ := jwt.ParseToken(c.Query("token"))
	user := models.GetUserById(claims.Id)

	var signedAgreement SignedAgreementRequest
	if err := c.ShouldBindJSON(&signedAgreement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(signedAgreement.Id) > 0 {
		gorm.GetDB().Exec("INSERT INTO user_agreement (user_id, agreement_id) values (?, ?)", user.ID, signedAgreement.Id)
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id": user.ID,
		"agreement_id": signedAgreement.Id,
	})
}