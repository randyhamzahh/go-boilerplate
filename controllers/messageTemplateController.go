package controllers

import (
	model "go-auth/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DBConnMessageTemplateController struct {
	DB *gorm.DB
}

func NewMessageTemplateController(db *gorm.DB) *DBConnMessageTemplateController {
	return &DBConnMessageTemplateController{
		DB: db,
	}
}

func (db *DBConnMessageTemplateController) StoreTemplate(c *gin.Context) {
	var body struct {
		Name    string
		Message string
		UserID  int
	}

	if c.Bind(&body) != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	template := model.MessageTemplate{
		Name:    body.Name,
		Message: body.Message,
		UserID:  body.UserID,
	}

	result := db.DB.Create(&template)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": template,
	})
}
