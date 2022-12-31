package server

import "github.com/gin-gonic/gin"

func StartServer() {
	router := gin.Default()

	registerApiRoutes(router)

	router.Run("0.0.0.0:9090")
}

func registerApiRoutes(router *gin.Engine) {

}
