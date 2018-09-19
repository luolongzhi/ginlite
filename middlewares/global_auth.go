package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

func GlobalAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		u := uuid.Must(uuid.NewV4())
		c.Set("_self_context__request_id", fmt.Sprintf("%v", u))

		c.Next()
	}
}
