package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func formatError(message string) map[string]string {
	return map[string]string{
		"message": message,
	}
}

func abortWithGenericError(ctx *gin.Context, err error) {
	if err == nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, formatError("Unknown error"))
		return
	}
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, formatError(err.Error()))
}

func abortWithBadRequest(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, formatError(err.Error()))
}

func abortWithMessage(ctx *gin.Context, status int, msg string) {
	ctx.AbortWithStatusJSON(status, formatError(msg))
}
