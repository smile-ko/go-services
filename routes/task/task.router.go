package task

import (
	c "go-services/internal/controllers"

	"github.com/gin-gonic/gin"
)

type TaskRouter struct {
}

func (t *TaskRouter) InitTaskRouter(router *gin.RouterGroup) {
	taskRouterPublic := router.Group("/task")

	{
		taskRouterPublic.GET("/:id", c.NewTaskController().GetOne)
		taskRouterPublic.GET("", c.NewTaskController().GetAll)
	}

	taskRouterPrivate := router.Group("/task")
	{
		taskRouterPrivate.POST("", c.NewTaskController().Create)
		taskRouterPrivate.PUT("/:id", c.NewTaskController().Update)
		taskRouterPrivate.DELETE("/:id", c.NewTaskController().Delete)
	}
}
