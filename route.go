package main

import (
    "github.com/gin-gonic/gin"
    "ginlite/controllers"
)

func InitRoute(r *gin.Engine) {
    ctrl := controllers.Init()

    r.GET("/ping", ctrl.TestController.Ping)
}

