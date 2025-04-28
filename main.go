package main

import (
	"go-auth/config"
	"go-auth/initializers"
	"go-auth/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	initializers.ConnectToPgSql()
}

func main() {

	router := gin.New()
	router.Use(gin.Logger())
	routes.RegisterAPIRoutes(router)

	router.GET("/ting", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "tong"})
	})

	router.Run(":" + config.AppConfig.AppPort)
}
