package initialize

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Run() (*gin.Engine, string) {
	// 0> Init zap logger
	logger := InitLogger()
	defer logger.Sync()

	// 1> Read config -> environment variables
	config, err := LoadConfig()
	if err != nil {
		logger.Fatal("Could not load config: %v", zap.Error(err))
	}

	// 3> Initialize database connection
	_, err = InitDB(&config)
	if err != nil {
		logger.Fatal("Could not initialize database: %v", zap.Error(err))
	}

	// 4> Initialize router
	r := InitRouter(config.LogLevel)

	// 5> Initialize other services if needed (e.g., cache, message queue, etc.)
	return r, config.ServerPort
}
