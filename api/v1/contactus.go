package apiv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	
	"goserver/libs/mail"
)

type Contact struct {
	Name    string `form:"name" json:"name" xml:"name"  binding:"required"`
	Email   string `form:"email" json:"email" xml:"email" binding:"required,email"`
	Website string `form:"website" json:"website" xml:"website" binding:"required,min=5"`
	Message string `form:"message" json:"message" xml:"message" binding:"required,min=20,max=2000"`
}

func ContactusApi(c *gin.Context) {
	var json Contact
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Thanks for contacting us. We will get back to you shortly.",
	})

	mail.SendContactUsEmail(json.Email, json.Name+" from "+json.Website+" has a message to us.", json.Message)
}