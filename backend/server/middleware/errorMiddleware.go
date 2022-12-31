package middleware

import "github.com/gin-gonic/gin"

func ApiErrorHandlerMw(ctx *gin.Context) {
	ctx.Next()

	var errMessageStack []string
	for _, err := range ctx.Errors {
		errMessageStack = append(errMessageStack, err.Error())
	}

	ctx.JSON(-1, errMessageStack)
}
