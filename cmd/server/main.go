package main

import (
	_ "github.com/leedev/go-rest-ddd/cmd/swag/docs"
	"github.com/leedev/go-rest-ddd/internal/initialize"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go Rest DDD API
// @version 1.0
// @description This is a sample server for a Go Rest DDD application.
// @host localhost:8800
// @BasePath /v1
func main() {
	r, port := initialize.Run()

	// prometheus.MustRegister(pingCounter)

	// r.GET("/ping/200", ping)
	// r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":" + port) // listen and serve on :8899
}
