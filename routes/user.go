package routes

import (
	"go-auth/controllers"
	"go-auth/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(routerGroup *gin.RouterGroup, userController *controllers.DBConnUserController) {
	user := routerGroup.Group("/user").Use(middleware.Authenticate())

	user.POST("/", userController.StoreUser)
	user.GET("/", userController.GetUsers)
	user.GET("/me", userController.GetMe)
	user.GET("/:id", userController.ShowUser)
	user.PUT("/:id", userController.UpdateUser)
	user.DELETE("/:id", userController.DeleteUser)
}
