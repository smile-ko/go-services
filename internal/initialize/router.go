package initialize

import (
	"go-services/global"
	"go-services/routes"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	userRouter := routes.RouterGroupApp.User
	taskRouter := routes.RouterGroupApp.Task

	mainRouter := r.Group("/api/v1")
	{
		userRouter.InitUserRouter(mainRouter)
		taskRouter.InitTaskRouter(mainRouter)
	}

	return r
}
