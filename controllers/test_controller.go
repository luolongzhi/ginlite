package controllers

import "github.com/gin-gonic/gin"

type TestController struct {
}

func (ctrl *TestController) Ping(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
