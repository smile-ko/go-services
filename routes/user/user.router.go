package user

import (
	c "go-services/internal/controllers"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(router *gin.RouterGroup) {

	userRouterPublic := router.Group("/user")
	{
		userRouterPublic.GET("/info", c.NewUserController().Index)
		userRouterPublic.GET("/:id", c.NewUserController().GetOne)
		userRouterPublic.GET("", c.NewUserController().GetAll)
	}

	userRouterPrivate := router.Group("/user")
	{
		userRouterPrivate.POST("", c.NewUserController().Create)
		userRouterPrivate.PUT("/:id", c.NewUserController().Update)
		userRouterPrivate.DELETE("/:id", c.NewUserController().Delete)
	}
}
