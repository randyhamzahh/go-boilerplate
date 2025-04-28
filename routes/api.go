package routes

import (
	"go-auth/controllers"
	"go-auth/initializers"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(router *gin.Engine) {
	userController := controllers.NewUserController(initializers.DB)
	authController := controllers.NewAuthController(initializers.DB)
	messageTemplate := controllers.NewMessageTemplateController(initializers.DB)

	api := router.Group("/api/v1")

	AuthRoutes(api, authController)
	UserRoutes(api, userController)
	MessageTemplateRoutes(api, messageTemplate)
}
