package initialize

import (
	"fmt"
	"go-services/global"
)

func Run() {
	LoadConfig()
	InitPostgresql()
	InitKafka()

	// Init router
	r := InitRouter()
	r.Run(fmt.Sprintf(":%d", global.Config.Server.Port))
}
