package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func AppAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
        fmt.Println("========> this is app auth")

		ctx.Next()
	}
}
