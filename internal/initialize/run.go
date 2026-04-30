package initialize

import (
	"hta-platform/global"

	"hta-platform/internal/middleware"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Run() (*gin.Engine, string) {
	// 0> Init zap logger
	logger := InitLogger()
	defer logger.Sync()

	// 1> Read config -> environment variables
	err := LoadConfig()
	if err != nil {
		logger.Fatal("Could not load config: %v", zap.Error(err))
	}

	// 2> Initialize Auth0 JWKS
	if global.ConfigValue.Auth0Domain != "" {
		if err := middleware.InitAuth0(); err != nil {
			logger.Fatal("Could not initialize Auth0", zap.Error(err))
		}
	}

	// 3> Initialize database connection
	db, err := InitDB()
	if err != nil {
		logger.Fatal("Could not initialize database: %v", zap.Error(err))
	}

	// 4> Initialize router
	r := InitRouter(db)

	// 5> Initialize other services if needed (e.g., cache, message queue, etc.)
	return r, global.ConfigValue.ServerPort
}
