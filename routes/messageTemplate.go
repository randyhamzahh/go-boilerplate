package routes

import (
	"go-auth/controllers"
	"go-auth/middleware"

	"github.com/gin-gonic/gin"
)

func MessageTemplateRoutes(routerGroup *gin.RouterGroup, messageTemplateController *controllers.DBConnMessageTemplateController) {

	route := routerGroup.Group("/template").Use(middleware.Authenticate())

	route.POST("/", messageTemplateController.StoreTemplate)
}
