package controllers

import ( 
    "github.com/gin-gonic/gin"
)

type TestController struct {
}

// Test Ping
// @Summary Test Ping 
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "We need ID!!"
// @Failure 404 {string} string  "Can not find ID"
// @Router /ping [get]
func (ctrl *TestController) Ping(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
