package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/config/v2"
	controller "rlesjak.com/ha-scheduler/server/controllers"
	"rlesjak.com/ha-scheduler/server/ui"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func StartServer(config config.Config) *gin.Engine {
	router := gin.Default()

	// Setup CORS
	router.Use(CORSMiddleware())

	ui.RegisterEmbeddedUiRoutes(router, config)
	registerApiRoutes(router, config)
	router.Run("0.0.0.0:9090")

	return router
}

func registerApiRoutes(router *gin.Engine, config config.Config) {
	apiConf := config.StringMap("Api")
	v1Path := apiConf["BasePath"] + apiConf["v1"]

	// Register Api error handler middleware
	// router.Use(middleware.ApiErrorHandlerMw)

	v1 := router.Group(v1Path)
	controller.RegisterElementGroupsController(v1)
	controller.RegisterTriggerGroupsController(v1)
	controller.RegisterTriggersController(v1)
}
