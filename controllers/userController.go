package controllers

import (
	model "go-auth/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DBConnUserController struct {
	DB *gorm.DB
}

func NewUserController(db *gorm.DB) *DBConnUserController {
	return &DBConnUserController{
		DB: db,
	}
}

func (u *DBConnUserController) GetMe(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (u *DBConnUserController) GetUsers(c *gin.Context) {

	var users []model.User

	result := u.DB.Find(&users)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (u *DBConnUserController) ShowUser(c *gin.Context) {
	var user model.User
	result := u.DB.First(&user, c.Param("id"))

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (u *DBConnUserController) StoreUser(c *gin.Context) {
	var body struct {
		Name  string
		Email string
	}

	if c.Bind(&body) != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	user := model.User{Name: body.Name, Email: body.Email}
	result := u.DB.Create(&user)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (u *DBConnUserController) UpdateUser(c *gin.Context) {
	var body struct {
		Name  string
		Email string
	}

	if c.Bind(&body) != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var user model.User
	u.DB.First(&user, c.Param("id"))

	result := u.DB.Model(&user).Updates(model.User{
		Name:  body.Name,
		Email: body.Email,
	})

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (u *DBConnUserController) DeleteUser(c *gin.Context) {
	var user model.User
	u.DB.First(&user, c.Param("id"))
	result := u.DB.Delete(&user)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"succes": true,
	})
}
