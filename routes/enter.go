package routes

import (
	"go-services/routes/task"
	"go-services/routes/user"
)

type RouterGroup struct {
	User user.UserRouterGroup
	Task task.TaskRouterGroup
}

var RouterGroupApp = new(RouterGroup)
