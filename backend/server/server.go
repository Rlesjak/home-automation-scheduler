package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/config/v2"
	controller "rlesjak.com/ha-scheduler/server/controllers"
)

func StartServer(config config.Config) *gin.Engine {
	router := gin.Default()

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
}
