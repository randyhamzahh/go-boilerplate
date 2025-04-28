package routes

import (
	"go-auth/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(routerGroup *gin.RouterGroup, authController *controllers.AuthController) {
	auth := routerGroup.Group("/")

	auth.POST("/signup", authController.RegisterUser)
	auth.POST("/login", authController.LoginUser)
}
