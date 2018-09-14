package controllers

import "github.com/gin-gonic/gin"

type TestController struct {
    Controller
}

func (ctrl *TestController) Ping(c *gin.Context) {
    ctrl.TestPrint()

    c.JSON(200, gin.H{
        "message": "pong",
    })
}



