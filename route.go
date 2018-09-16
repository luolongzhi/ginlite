package main

import (
    "github.com/gin-gonic/gin"
    "ginlite/controllers"
)

func RouteRegister(r *gin.Engine) {
    r.GET("/ping", (&controllers.TestController{}).Ping)
    r.POST("/users", (&controllers.UserController{}).Create)
    r.POST("/users.delete", (&controllers.UserController{}).Delete)
}

