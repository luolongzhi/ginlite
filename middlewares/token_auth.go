package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func TokenAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
        fmt.Println("========> this is token auth")

		ctx.Next()
	}
}
