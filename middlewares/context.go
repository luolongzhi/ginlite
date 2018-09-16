package middlewares

import "github.com/gin-gonic/gin"

func SetContext(key string, value interface{}) gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Set(key, value)
        c.Next()
    }
}
