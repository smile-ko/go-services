package user

import (
	"go-services/internal/controllers"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	userRouterPublic := router.Group("/user")
	{
		userController := controllers.NewUserController()
		userRouterPublic.GET("/info", userController.Index)
		userRouterPublic.POST("", userController.Create)
		userRouterPublic.GET("/:id", userController.GetOne)
		userRouterPublic.PUT("/:id", userController.Update)
		userRouterPublic.DELETE("/:id", userController.Delete)
	}
}
